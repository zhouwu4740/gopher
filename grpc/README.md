# gRPC

## Install
- protoc
- go install google.golang.org/protobuf/cmd/protoc-gen-go
- go install google.golang.org/protobuf/cmd/protoc-gen-go-grpc
- go install google.golang.org/grpc

### Quick Start
1. 编写 proto 协议文件
```grpc
syntax = "proto3";
option go_package = "user/hello";

message Request {
  string name = 1;
}

message Response {
  string message = 2;
}

service Greeter {
  rpc Hello(Request) returns (Response) {}
}
```
2. 生成代码
```shell
protoc -I $GOPATH/src  -I . --go_out=.  --go-grpc_out=.  hello.proto
```
3. 创建 struct 并实现GreeterServer interface，并将struct 注册到gRPC server