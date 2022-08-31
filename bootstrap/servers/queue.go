package servers

import (
	"context"
	"encoding/json"
	app2 "github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/app/queues"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/app"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

// Queue @Bean("queue")
type Queue struct {
	// 队列配置文件的所有配置
	fileConfig *services.Config `inject:"config, queue"`
	// 队列具体配置
	queueConfig *services.Config `inject:"config, queue.queue"`
	// 连接
	Connect *services.Redis `inject:"database, @config(queue.connection)"`

	// 执行限速度
	limit     uint // 如果 == 0 , 则停止
	limitChan chan bool
	// 所有路由注册
	route map[string]constraint.Job
	// 支持并发调用job, 这个是route副本保存结构
	dispatch sync.Map

	// 打开广播进程, 允许把某个message投递到广播队列, 所有服务都会收到
	broadcast  *Broadcast
	delayQueue DelayQueue
}

func (q *Queue) Init() {
	q.route = make(map[string]constraint.Job)
	q.limit = uint(q.queueConfig.GetInt("worker_limit", 100))
	q.limitChan = make(chan bool, q.limit)

	q.Listen(queues.GetAllProvider())
}

func (q *Queue) StartBroadcast() {
	q.broadcast = &Broadcast{
		Connect: q.Connect,
		queue:   q,
	}
}

func (q *Queue) CloseBroadcast() {
	q.broadcast.Close()
}

func (q *Queue) HasBroadcast() bool {
	return q.broadcast != nil
}

// Push 投递队列
func (q *Queue) Push(message interface{}) {
	jsonStr, err := json.Marshal(message)
	if err != nil {
		log.Error(err)
		return
	}
	stream, _ := q.getJobInfo(message)
	route := jobToRoute(message)
	d := q.Connect.Client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: stream,
		MaxLen: int64(q.queueConfig.GetInt("stream_limit", 10000)),
		Values: map[string]interface{}{
			"route": route,
			"event": jsonStr,
		},
	})
	if d.Err() != nil {
		log.Error("Queue push, redis xadd error: ", d.Err())
	}
	q.Delay(time.Second)
}

// Delay 延后投递执行
func (q *Queue) Delay(t time.Duration) *DelayTask {
	return &DelayTask{
		interval: t,
		queue:    q,
		message:  nil,
	}
}

// Publish 对message广播
func (q *Queue) Publish(message interface{}, topics ...string) {
	jsonStr, err := json.Marshal(message)
	if err != nil {
		log.Error(err)
		return
	}

	if len(topics) == 0 {
		topics = append(topics, app2.Config("queue.broadcast.topic", "home_broadcast"))
	}

	route := jobToRoute(message)
	for _, topic := range topics {
		q.broadcast.Publish(topic, Msg{
			Route: route,
			Data:  string(jsonStr),
		})
	}
}

func (q *Queue) Run() {
	q.initStream()

	for route, job := range q.route {
		q.dispatch.Store(route, reflect.New(reflect.TypeOf(job).Elem()))
	}
	serialQueue := make([]interface{}, 0)
	baseQueue := make([]interface{}, 0)
	for _, job := range q.route {
		s, ok := job.(constraint.SetSerial)
		if ok && s.SetSerial() {
			serialQueue = append(serialQueue, job)
		} else {
			baseQueue = append(baseQueue, job)
		}
	}
	q.runBaseQueueList(baseQueue)

	// 必须串行执行的job
	if len(serialQueue) != 0 {
		if app.HasBean("election") {
			// 在选举服务之后启动
			app.GetBean("election").(app.AppendRun).AppendRun(func() {
				q.runSerialQueueList(serialQueue)
			})
		} else {
			log.Info("存在串行执行的job, 但没有注册 election 服务, 如果有多个副本运行可能得不到预想效果")
			q.runSerialQueueList(serialQueue)
		}
	}

	// 广播进程
	if q.HasBroadcast() {
		go q.broadcast.Subscribe(q.fileConfig.GetString("broadcast.topic", "home_broadcast"))
	}

	if q.delayQueue != nil {
		q.delayQueue.Run()
	}
}

func (q *Queue) runBaseQueueList(list []interface{}) {
	groupQueue := make(map[string]map[string]string)
	for _, job := range list {
		stream, group := q.getJobInfo(job)
		if groupQueue[group] == nil {
			groupQueue[group] = make(map[string]string)
		}
		groupQueue[group][stream] = ">"
	}
	for group, streams := range groupQueue {
		ruS := make([]string, 0)
		for s, s2 := range streams {
			ruS = append(ruS, s, s2)
		}
		go q.runBaseQueue(group, ruS)
	}
}

func (q *Queue) runBaseQueue(group string, streams []string) {
	ctx := context.Background()
	Hostname, _ := os.Hostname()
	consumer := strings.ReplaceAll(q.queueConfig.GetString("consumer_name"), "{hostname}", Hostname)
	block := time.Duration(q.queueConfig.GetInt("stream_block", 60)) * time.Second
	for q.limit > 0 {
		cmd := q.Connect.Client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  streams,
			Count:    1,
			Block:    block,
			NoAck:    false,
		})

		if cmd.Err() == nil {
			result, _ := cmd.Result()
			for _, XStream := range result {
				for _, xMessage := range XStream.Messages {
					route, ok := xMessage.Values["route"]
					if !ok {
						log.Error("队列任务数据无法解析route, 一个job的结构必须是 {event:'', route:''}")
						continue
					}
					event, ok := xMessage.Values["event"]
					if !ok {
						log.Error("队列任务数据无法解析event, 一个job的结构必须是 {event:'', route:''}")
						continue
					}

					job, ok := q.dispatch.Load(route.(string))
					if !ok {
						log.Errorf("无法处理的route: %v", route.(string))
						continue
					}
					go q.runJob(job.(reflect.Value), event.(string), xMessage.ID, XStream.Stream, group)
					q.limitChan <- true
				}
			}
		} else if cmd.Err().Error() == "redis: nil" {
			time.Sleep(3 * time.Second)
		} else {
			log.Errorf("redis命令未知错误, XReadGroup, %v", cmd.Err().Error())
			time.Sleep(60 * time.Second)
		}
	}
}

func (q *Queue) runSerialQueueList(list []interface{}) {
	groupQueue := make(map[string]map[string]string)
	for _, job := range list {
		stream, group := q.getJobInfo(job)
		if groupQueue[group] == nil {
			groupQueue[group] = make(map[string]string)
		}
		groupQueue[group][stream] = ">"
	}
	for group, streams := range groupQueue {
		ruS := make([]string, 0)
		for s, s2 := range streams {
			ruS = append(ruS, s, s2)
		}
		go q.runSerialQueue(group, ruS)
	}
}

func (q *Queue) runSerialQueue(group string, streams []string) {
	ctx := context.Background()
	Hostname, _ := os.Hostname()
	consumer := strings.ReplaceAll(q.queueConfig.GetString("consumer_name"), "{hostname}", Hostname)
	block := time.Duration(q.queueConfig.GetInt("stream_block", 60)) * time.Second
	limitChan := make(chan bool, 1)
	for q.limit > 0 {
		cmd := q.Connect.Client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  streams,
			Count:    1,
			Block:    block,
			NoAck:    false,
		})

		if cmd.Err() == nil {
			result, _ := cmd.Result()
			for _, XStream := range result {
				for _, xMessage := range XStream.Messages {
					route, ok := xMessage.Values["route"]
					if !ok {
						log.Error("队列任务数据无法解析route, 一个job的结构必须是 {event:'', route:''}")
						continue
					}
					event, ok := xMessage.Values["event"]
					if !ok {
						log.Error("队列任务数据无法解析event, 一个job的结构必须是 {event:'', route:''}")
						continue
					}

					job, ok := q.dispatch.Load(route.(string))
					if !ok {
						log.Errorf("无法处理的route: %v", route.(string))
						continue
					}
					go func(job reflect.Value, event string, id, stream, group string) {
						defer func() {
							if err := recover(); err != nil {
								log.Error("队列任务执行发生错误", err)
							}
							<-limitChan
						}()
						v := job.Interface()
						newJob, ok := v.(constraint.Job)
						if ok {
							err := json.Unmarshal([]byte(event), newJob)
							if err == nil {
								newJob.Handler()
								q.Connect.Client.XAck(context.Background(), group, stream, id)
							} else {
								log.Errorf("runJob, json.Unmarshal data err = %v", err)
							}
						}
					}(job.(reflect.Value), event.(string), xMessage.ID, XStream.Stream, group)
					limitChan <- true
				}
			}
		} else if cmd.Err().Error() == "redis: nil" {
			time.Sleep(3 * time.Second)
		} else {
			log.Error(cmd.Err().Error())
			time.Sleep(60 * time.Second)
		}
	}
}

func (q *Queue) Exit() {
	q.limit = 0
}

func (q *Queue) Listen(jobs []interface{}) {
	// 注册路由绑定Job
	for _, job := range jobs {
		handle, ok := job.(constraint.Job)
		if ok {
			// job 绑定的路由
			route := ""
			mr, ok := job.(constraint.SetRoute)
			if ok {
				route = mr.SetRoute()
			} else {
				// route = message.*
				route = jobToRoute(handle)
			}
			q.AddJob(route, handle)
		}
	}
}

func (q *Queue) AddJob(route string, handle constraint.Job) {
	if _, o := q.route[route]; !o {
		q.route[route] = handle
	} else {
		panic("队列路由: " + route + " 重复, 需要在您的message对象创建 SetRoute() string , 使用自定义路由避免重复。")
	}
}

// 自动计算struct路径
func jobToRoute(handle interface{}) string {
	mr, ok := handle.(constraint.SetRoute)
	if ok {
		return mr.SetRoute()
	}

	ref := reflect.TypeOf(handle)
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	// message
	mty := ref.PkgPath() + ref.Name()
	if strings.Index(mty, "message") != -1 {
		return mty
	}
	// job
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		ty := field.Type.PkgPath() + field.Type.Name()
		if strings.Index(ty, "message") != -1 {
			return ty
		}
	}
	panic("自动注册路由的信息总线的对象路径必须包含 `message`。")
}

func (q *Queue) initStream() {
	ctx := context.Background()
	streams := make(map[string]bool)
	for _, job := range q.route {
		stream, group := q.getJobInfo(job)
		if !streams[stream] {
			continue
		}
		xInfoG, err := q.Connect.Client.XInfoGroups(ctx, stream).Result()
		streams[stream] = true
		if err != nil {
			if err.Error() != "ERR no such key" {
				log.Warn(err)
				continue
			} else {
				q.Connect.Client.XAdd(ctx, &redis.XAddArgs{
					Stream: stream,
					ID:     "1",
					Values: map[string]interface{}{
						"route": "",
						"event": "init",
					},
				})
				q.Connect.Client.XDel(ctx, stream, "1")
			}
		}
		mInfoG := make(map[string]redis.XInfoGroup)
		for _, g := range xInfoG {
			mInfoG[g.Name] = g
		}
		if _, ok := mInfoG[group]; !ok {
			q.Connect.Client.XGroupCreate(ctx, stream, group, "$")
		}
	}
}

func (q *Queue) getJobInfo(job interface{}) (stream string, group string) {
	jobStream, ok := job.(constraint.SetQueue)
	if ok {
		stream = jobStream.SetQueue()
	} else {
		stream = q.queueConfig.GetString("stream_name", "home_default_stream")
	}

	jobGroup, ok := job.(constraint.SetGroup)
	if ok {
		group = jobGroup.SetGroup()
	} else {
		group = q.queueConfig.GetString("group_name", "home_default_group")
	}
	return stream, group
}

func (q *Queue) jobToGroup(job interface{}) string {
	g, ok := job.(constraint.SetGroup)
	if ok {
		return g.SetGroup()
	}
	return q.queueConfig.GetString("group_name", "home_default_group")
}

func (q *Queue) runJob(job reflect.Value, event string, id, stream, group string) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("队列任务执行发生错误", err)
		}

		<-q.limitChan
	}()
	v := job.Interface()
	newJob, ok := v.(constraint.Job)
	if ok {
		err := json.Unmarshal([]byte(event), newJob)
		if err == nil {
			newJob.Handler()
			q.Connect.Client.XAck(
				context.Background(),
				group,
				stream,
				id,
			)
		} else {
			log.Errorf("runJob, json.Unmarshal data err = %v", err)
		}
	}
}

// StartDelayQueue 开启延时队列
func (q *Queue) StartDelayQueue() {
	if app.HasBean("delay_queue") {
		q.delayQueue = app.GetBean("delay_queue").(DelayQueue)
	} else {
		q.delayQueue = NewDelayQueueForMysql()
	}
}

// DelayTask 延时队列包装
type DelayTask struct {
	// 延后多少时间投递
	interval time.Duration
	queue    *Queue
	message  interface{}
}

func (d *DelayTask) Push(message interface{}) {
	d.message = message

	d.queue.delayQueue.Push(*d)
}
