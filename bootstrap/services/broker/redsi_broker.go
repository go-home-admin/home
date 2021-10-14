package broker

import (
	"context"
	"encoding/json"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"reflect"
	"time"
)

type Config struct {
	Stream     string
	NoMkStream bool
	MaxLen     int64 // MAXLEN N
	// Approx causes MaxLen and MinID to use "~" matcher (instead of "=").
	Approx bool
	Limit  int64
}

// @Bean
type RedisBroker struct {
	client      *redis.Client
	config      map[string]*Config
	dispatch    map[string]constraint.Job
	streamGroup map[string][]string
}

func (r *RedisBroker) Init() {
	r.config = make(map[string]*Config)
	r.dispatch = make(map[string]constraint.Job)
	r.streamGroup = make(map[string][]string)
}

func defConfig() Config {
	return Config{
		Stream:     constraint.QueueName,
		NoMkStream: false,
		MaxLen:     100000,
		Approx:     true,
		Limit:      0,
	}
}

func (r *RedisBroker) SetConfig(client constraint.RedisClient, config ...Config) {
	r.client = client.GetClient()

	if len(config) == 0 {
		config = append(config, defConfig())
	}
	for _, c := range config {
		r.config[c.Stream] = &c
		r.streamGroup[c.Stream] = make([]string, 0)
	}
}

func (r *RedisBroker) Push(event interface{}, names ...string) {
	if len(names) == 0 {
		names = append(names, constraint.QueueName)
	}
	for _, name := range names {
		c, ok := r.config[name]
		if ok {
			ref := reflect.TypeOf(event)
			eventKey := ref.PkgPath() + ref.Name()
			jsonStr, err := json.Marshal(event)
			if err != nil {
				log.Error(err)
				continue
			}
			ret := r.client.XAdd(context.Background(), &redis.XAddArgs{
				Stream:     c.Stream,
				NoMkStream: false,
				MaxLen:     c.MaxLen,
				Approx:     false,
				Limit:      c.Limit,
				Values: map[string]interface{}{
					"route": eventKey,
					"event": jsonStr,
				},
			})
			if ret.Err() != nil {
				log.Error(ret.Err())
			}
		}
	}
}

func (r *RedisBroker) Consumer(job constraint.Job) {
	c := job.Config()
	event := c.Event
	group := c.GroupName
	queue := c.QueueName

	if group == "" {
		group = constraint.QueueGroup
	}
	if queue == "" {
		queue = constraint.QueueName
	}

	ref := reflect.TypeOf(event)
	eventKey := ref.PkgPath() + ref.Name()
	// 调度器
	r.dispatch[eventKey] = job

	_, ok := r.streamGroup[queue]
	if !ok {
		log.Warning("消费者分组, 没有初始化, 请在调用SetConfig时, 指定第二个参数")
		r.streamGroup[queue] = make([]string, 0)
		temp := defConfig()
		r.config[queue] = &temp
	}
	// 消费者分组
	r.streamGroup[queue] = append(r.streamGroup[queue], group)
}

// Loop 开启消费已经设置的stream
func (r *RedisBroker) Loop() {
	ctx := context.Background()

	for queueName, queueGroups := range r.streamGroup {
		xInfoG, err := r.client.XInfoGroups(ctx, queueName).Result()
		if err != nil {
			continue
		}
		mInfoG := make(map[string]redis.XInfoGroup)
		for _, group := range xInfoG {
			mInfoG[group.Name] = group
		}
		for _, queueGroup := range queueGroups {
			// 创建消费组
			if _, ok := mInfoG[queueGroup]; !ok {
				r.client.XGroupCreate(ctx, queueName, queueGroup, "$")
			}

			go r.read(queueGroup, queueName)
		}
	}
}
func (r *RedisBroker) read(group string, queueName string) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("队列工人发生错误, 已退出", err)
		}
	}()
	ctx := context.Background()
	for {
		cmd := r.client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: "",
			Streams:  []string{queueName, ">"},
			Count:    1,
			Block:    3 * time.Second,
			NoAck:    false,
		})

		if cmd.Err() == nil {
			result, _ := cmd.Result()
			for _, XStream := range result {
				for _, xMessage := range XStream.Messages {
					route, ok := xMessage.Values["route"]
					if !ok {
						log.Error("队列任务数据无法解析route")
						continue
					}
					event, ok := xMessage.Values["event"]
					if !ok {
						log.Error("队列任务数据无法解析event")
						continue
					}

					job := r.dispatch[route.(string)]
					err := json.Unmarshal([]byte(event.(string)), job)
					if err != nil {
						log.Error("无法解析任务数据", err)
						continue
					}

					task := constraint.Task{
						ID:    xMessage.ID,
						Event: job,
					}

					job.Handler(task)
					// 自动ack
					r.client.XAck(ctx, queueName, group, xMessage.ID)
				}
			}
		}
	}
}
