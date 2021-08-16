# home
文档先行, 舒适、工具化、开箱即用的、统一编程风格框架; 系统地集成各种热门组件到框架中, 对web后端业务开发封装最佳实践。

# 安装
建议把代码和工具安装到同一个父级目录下
~~~~shell
cd go-home-admin
git clone git@github.com:go-home-admin/home.git
git clone git@github.com:go-home-admin/home-toolset-php.git
~~~~
#### 代码生成辅助工具需要初始化
~~~~shell
cd home-toolset-php
composer install
~~~~
#### 启动`home`
~~~~shell
cd home
make dev
~~~~

# 其他
### 国内代理设置
~~~~shell
go env -w GOPROXY=https://goproxy.io,direct
~~~~

### GoLand 编辑器辅助设置


# TODO
- [ ] 框架引导
- [ ] 基础http server
- [ ] proto生成api路由swagger和基础代码, grpc入口
- [ ] 队列服务
- [ ] 定时服务