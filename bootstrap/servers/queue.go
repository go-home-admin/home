package servers

import (
	"github.com/go-home-admin/home/bootstrap/constraint"
	"github.com/go-home-admin/home/bootstrap/providers"
	"github.com/go-home-admin/home/bootstrap/services"
	"github.com/go-redis/redis/v8"
	"reflect"
	"strings"
)

// Queue @Bean
type Queue struct {
	// 队列配置文件的所有配置
	queue *services.Config `inject:"config, queue"`
	// 队列具体配置
	config *services.Config
	// 连接
	Connect *redis.Client

	route         map[string]constraint.Job
	queueMapGroup map[string]string
}

func (q *Queue) Init() {
	q.route = make(map[string]constraint.Job)
	q.queueMapGroup = make(map[string]string)
	q.config = services.NewConfig(q.queue.GetKey("queue"))
	q.Connect = providers.NewRedisProvider().GetBean(q.queue.GetString("connection")).(*redis.Client)
}

func (q *Queue) Run() {

}

func (q *Queue) Exit() {}

func (q *Queue) Listen(jobs []interface{}) {
	// 注册路由绑定Job
	for _, job := range jobs {
		handle, ok := job.(constraint.Job)
		if ok {
			route := ""
			mr, ok := job.(constraint.SetRoute)
			if ok {
				route = mr.SetRoute()
			} else {
				// route = message.*
				route = jobToRoute(handle)
			}
			if _, o := q.route[route]; !o {
				q.route[route] = handle
			} else {
				panic("队列路由: " + route + " 重复, 需要在您的message对象创建 SetRoute() string , 使用自定义路由避免重复。")
			}
		}

	}
}

func (q *Queue) Push(message interface{}) {

}

func jobToRoute(handle interface{}) string {
	ref := reflect.TypeOf(handle)
	ref = ref.Elem()

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		ty := field.Type.String()
		if strings.Index(ty, "message.") != -1 {
			return ty
		}
	}
	panic("自动注册路由的信息总线的对象路径必须包含 `message.`。")
}
