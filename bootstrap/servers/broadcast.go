package servers

import (
	"context"
	"encoding/json"
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"reflect"
)

type Broadcast struct {
	Connect *services.Redis
	queue   *Queue
	Run     bool
}

func (b *Broadcast) Close() {
	b.Run = false
}

// Subscribe 订阅主题
func (b *Broadcast) Subscribe(topic string) {
	b.Run = true
	for b.Run {
		b.subscribeFor(topic)
	}
}

func (b *Broadcast) subscribeFor(topic string) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Subscribe 发生错误", err)
		}
	}()

	pubSub := b.Connect.Client.Subscribe(context.Background(), topic)
	defer func(pubSub *redis.PubSub) {
		err := pubSub.Close()
		if err != nil {
			log.Error(err)
		}
	}(pubSub)
	if _, err := pubSub.Receive(context.Background()); err != nil {
		log.Error("failed to receive from control PubSub", err)
		return
	}

	for b.Run {
		for msg := range pubSub.Channel() {
			m := &Msg{}
			err := json.Unmarshal([]byte(msg.Payload), m)
			if err != nil {
				log.Errorf("广播队列无法处理的msg: %v", msg)
				continue
			}

			job, ok := b.queue.dispatch.Load(m.Route)
			if !ok {
				log.Errorf("广播队列无法处理的route: %v", m)
				continue
			}

			go b.runJob(job.(reflect.Value), m.Data)
		}
	}
}

func (b *Broadcast) Publish(topic string, m Msg) {
	msg, err := json.Marshal(m)
	if err != nil {
		log.Error(err)
		return
	}
	b.Connect.Client.Publish(context.Background(), topic, msg)
}

func (b *Broadcast) runJob(job reflect.Value, event string) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("广播队列任务执行发生错误", err)
		}
	}()
	v := job.Interface()
	newJob, ok := v.(constraint.Job)
	if ok {
		err := json.Unmarshal([]byte(event), newJob)
		if err == nil {
			newJob.Handler()
		} else {
			log.Errorf("broadcast runJob, json.Unmarshal data err = %v", err)
		}
	}
}

type Msg struct {
	Route string `json:"route"`
	Data  string `json:"event"`
}
