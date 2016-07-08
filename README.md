#skyintface

## 1. 加载protobuf

   protoc-gen-go:

   source profile

```shell
go get github.com/golang/protobuf/protoc-gen-go

cd github.com/golang/protobuf/protoc-gen-go

go build

go install

vi /etc/profile 将$GOPATH/bin #加入环境变量
source profile
```



​    proto:

```shell
go get github.com/golang/protobuf/proto
cd github.com/golang/protobuf/proto
go build
go install
```



## 2.构建主题程序

​    go build main

## 3.构建模拟客户端

​    go build client