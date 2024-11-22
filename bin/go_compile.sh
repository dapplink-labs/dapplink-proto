#!/bin/bash

# 检查是否安装了 Go
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed" >&2
    exit 1
fi

# 确保 GOPATH 已设置
if [ -z "$GOPATH" ]; then
    export GOPATH=$HOME/go
    echo "GOPATH was not set, using default: $GOPATH"
fi

# 创建必要的目录
mkdir -p $GOPATH/bin

# 安装必要的 protoc 插件
echo "Installing required Go plugins..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
if [ $? -ne 0 ]; then
    echo "Failed to install protoc-gen-go" >&2
    exit 1
fi

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
if [ $? -ne 0 ]; then
    echo "Failed to install protoc-gen-go-grpc" >&2
    exit 1
fi

# 添加 GOBIN 到 PATH
export PATH=$PATH:$GOPATH/bin

echo "Compiling protobuf files..."
protoc -I ./ \
    --go_out=./ \
    --go-grpc_out=require_unimplemented_servers=false:. \
    dapplink/*.proto

if [ $? -eq 0 ]; then
    echo "Compilation completed successfully"
else
    echo "Compilation failed" >&2
    exit 1
fi
