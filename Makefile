GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# 安装代码工具(开发机器需要)
# export GOPATH=$HOME/go PATH=$PATH:$GOPATH/bin
install:
	brew install protobuf
	go get google.golang.org/grpc						# 原始微服务工具
	go get -u github.com/golang/protobuf/proto			# proto 工具链
	go get -u github.com/golang/protobuf/protoc-gen-go	# proto 工具链

# 只维护 protoc
protoc:
	php bin/toolset protoc

make-route:
	php bin/toolset make:route

make-bean:
	php bin/toolset make:bean ./

# 调试启动
dev:protoc make-bean
	go run main.go --path=config.ini

