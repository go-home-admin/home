GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# 安装代码工具(开发机器需要)
# export GOPATH=$HOME/go PATH=$PATH:$GOPATH/bin
mac-install:
	brew install protobuf								# mac下自动安装, win环境手动安装
	go install google.golang.org/grpc						# 原始微服务工具
	go install github.com/golang/protobuf/proto			# proto 工具链
	go install github.com/golang/protobuf/protoc-gen-go	# proto 工具链, 生成go代码插件

# Orm自动维护
make-orm:
	go run ./bin/toolset/main.go make:protoc

# 只维护 protoc
protoc:
	go run ./bin/toolset/main.go make:protoc

make-route:
	go run ./bin/toolset/main.go make:route

make-swagger:
	go run ./bin/toolset/main.go make:protoc

make-bean:
	go run ./bin/toolset/main.go make:bean

# 调试启动
dev:protoc make-route make-bean
	go run main.go --path=./config

