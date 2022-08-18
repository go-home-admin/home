GOBIN := $(shell go env GOBIN)
ATDIR := $(shell pwd)

# 安装代码工具(开发机器需要)
# export GOPATH=$HOME/go PATH=$PATH:$GOPATH/bin
mac-install:
	protoc --version || brew install protobuf					# mac下自动安装, win环境手动安装
	cd ~ && go install github.com/golang/protobuf/proto			# proto 工具链
	cd ~ && go install github.com/golang/protobuf/protoc-gen-go	# proto 工具链, 生成go代码插件
	cd ~ && go install github.com/go-home-admin/toolset@latest

make-bean:
	toolset  make:bean

