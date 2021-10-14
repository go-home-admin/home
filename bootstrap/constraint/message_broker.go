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
	Event     interface{}
	GroupName string
	QueueName string
}

type Job interface {
	Config() JobConfig
	Handler(task Task)
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
