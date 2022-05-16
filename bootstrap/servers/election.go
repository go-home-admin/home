package servers

import (
	"context"
	"github.com/go-home-admin/home/app/message"
	"github.com/go-home-admin/home/app/queues"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-home-admin/home/bootstrap/services/app"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

// Election @Bean("election")
type Election struct {
	// 队列配置文件的所有配置
	*services.Config `inject:"config, election"`
	// 连接
	Connect   *services.Redis
	runUid    string // 进程唯一id
	key       string // 队列名称
	isRunNode bool
	lockTime  int
	awakens   []interface{}
}

func (k *Election) AppendRun(fun func()) {
	k.awakens = append(k.awakens, fun)
}

func (k *Election) Init() {
	k.awakens = make([]interface{}, 0)
	k.runUid = uuid.NewV4().String()
	k.key = k.GetString("default.key", "home_default_election")
	k.lockTime = k.GetInt("default.lock_time", 60)

	k.Connect = providers.NewRedisProvider().GetBean(k.GetString("connection", "redis")).(*services.Redis)
}

func (k *Election) Run() {
	if app.HasBean("queue") {
		job := queues.NewElectionClose()
		NewQueue().AddJob(jobToRoute(job), job)
	}
	// 标识当前节点抢到执行权利
	for !k.check() {
		time.Sleep(time.Duration(k.lockTime) * time.Second)
	}
	// 执行节点才可以走下去
	for _, awaken := range k.awakens {
		switch awaken.(type) {
		case func():
			awaken.(func())()
		case constraint.KernelServer:
			awaken.(constraint.KernelServer).Run()
		default:
			log.Warning("选举后才启动的参数, 请传入闭包或者constraint.KernelServer")
		}
	}
}

func (k *Election) Exit() {
	k.Connect.Client.Del(context.Background(), k.key)

	if app.HasBean("queue") {
		fileConfig := providers.GetBean("config").(providers.Bean).GetBean("queue").(*services.Config)
		topic := fileConfig.GetString("broadcast.topic", "home_broadcast")
		// 广播到其他副本
		NewQueue().Publish(topic, message.ElectionClose{
			Time: time.Now().Unix(),
		})
	}
}

func (k *Election) check() bool {
	if k.isRunNode {
		return true
	}

	// 沉默节点, 尝试检查runNode是否死机, 抢夺执行权利
	ctx := context.Background()
	ok := k.Connect.Client.SetNX(ctx, k.key, k.runUid, time.Duration(k.lockTime+10)*time.Second)

	if ok.Val() {
		k.isRunNode = true
		// 设置一个保持心跳的循环
		go func() {
			ex := time.Duration(k.lockTime) * time.Second
			for range time.Tick(ex) {
				k.Connect.Client.Expire(context.Background(), k.key, time.Duration(k.lockTime+10)*time.Second)
			}
		}()
		machine, _ := os.Hostname()
		log.Info(machine + " 通过了选举")
		return true
	} else {
		k.isRunNode = false
		return false
	}
}

// GetServer 提供统一命名规范的独立服务
func GetServer(leaders []interface{}) constraint.KernelServer {
	server := NewElection()

	server.awakens = append(server.awakens, leaders...)
	return server
}
