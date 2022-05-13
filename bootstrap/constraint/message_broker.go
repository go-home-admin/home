// 信息总线相关定义
package constraint

import "github.com/go-redis/redis/v8"

var QueueName = "home_stream"
var QueueGroup = "home_stream_group"

// 信息中间人
type Broker interface {
	Push(event interface{}, names ...string)
	Consumer(job Job)
	Loop()
}

type RedisClient interface {
	GetClient() *redis.Client
}

type JobConfig struct {
	Message   interface{}
	GroupName string
	QueueName string
}

type Job interface {
	Handler()
}

type SetRoute interface {
	SetRoute() string
}

type SetQueue interface {
	SetQueue() string
}

type SetGroup interface {
	SetGroup() string
}

// SetSerial 设置为穿行消费
type SetSerial interface {
	SetSerial() bool
}

type Task struct {
	ID    string
	Group string
	Queue string
	Event interface{}
}

type Worker interface {
	Run(job Job, task Task)
}
