# home
文档先行, 舒适、工具化、开箱即用的、统一编程风格框架; 系统地集成各种热门组件到框架中, 对web后端业务开发封装最佳实践。

# 安装
建议把代码和工具安装到同一个父级目录下
~~~~shell
cd go-home-admin
git clone https://github.com/go-home-admin/home.git
git clone https://github.com/go-home-admin/home-toolset-php.git
~~~~
#### 代码生成辅助工具需要初始化
~~~~shell
cd home-toolset-php
composer install
~~~~
#### 启动`home`, 需要检查依赖（protobuf, go, protoc-gen-go(protoc-gen-go需要依赖GOBIN环境变量)）
~~~~shell
cd home
make mac-install
make dev
~~~~

# 入门
### Makefile 工具
项目根目录的Makefile文件编写了所有的常用命令和流程, 启动`make dev`是保证能一键启动的命令(不包括环境处理)

### 添加新的模块和api
在home/protobuf创建新的目录admin, 创建proto文件, 同时设置和目录一致的package和option
~~~~go
package admin;
option go_package = "github.com/go-home-admin/home/generate/proto/admin";
service Public {
  option (http.RouteGroup) = "admin-public";

  rpc Index(IndexRequest)returns(IndexResponse){
    option (http.Get) = "/hello";
  }
}
~~~~
在这里app/http/kernel.go, 需要注册你的业务前缀, 中间件。

执行`make dev`就会生成文档和基础代码, 这时项目已经正启动和访问, 当然访问api只响应基础字段。

### 引导和依赖注入
在任意地方声明方法集合时, 加上@Bean就会被框架自动实例
~~~~go
// @Bean
type Controller struct {
    // logic 会被自动注入, 方法可以直接使用
    logic *logic.Demo `inject:""`
}
func (receiver *Controller) Test() {
    logic.Call()
}
~~~~
如果是在测试文件时, 不在引导流程的的结构体, 也可以正常获得并调用, InitializeNew{你的结构体}Provider是工具生成的
~~~~go
func Test_Login(t *testing.T) {
    InitializeNewControllerProvider().Test()
}
~~~~

### 注册自己的服务
~~~~go
// @Bean
type Sms struct {}
~~~~
只要任何结构体的字段是Sms类型, 并且使用了inject标签, Sms就可以正常实例化和使用
~~~~
// @Bean
type SmsProviders struct {
    sms *Sms `inject:""`
}
// 如果实现了Init()函数, 还会被自动调用
func (i *SmsProviders) Init() {}
~~~~

# 其他
### 国内代理设置
~~~~shell
go env -w GOPROXY=https://goproxy.io,direct
~~~~

### GoLand 编辑器辅助设置


# TODO
- [ ] Gorm
- [ ] Redis
- [ ] Swagger
- [ ] 队列服务
- [ ] 定时服务