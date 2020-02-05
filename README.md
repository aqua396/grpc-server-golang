# grpc-server-golang
tutorial and template [gRPC](https://grpc.io/) project

# How to start
## install go
1. install with [homebrew](https://brew.sh/index_ja).
```shell
$ brew install golang
```
2. set environment variable.
```shell
$ echo 'export GOPATH=$HOME/go' >> ~/.bash_profile
$ echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
$ source ~/.bash_profile
```
3. install [protocol buffers](https://developers.google.com/protocol-buffers).
```shell
$ brew install protobuf
```
4. get [protoc-gen-go](https://godoc.org/github.com/golang/protobuf/protoc-gen-go).
```shell
$ go get -u github.com/golang/protobuf/protoc-gen-go
```
5. create ```.proto``` file for service.
```golang
syntax = "proto3";
package greet;
// service interface
service Greet {
    rpc Say(Question) returns (Reply) {}
}
// request object
message Question {
    string message = 1;
}
// response object
message Reply {
    string message = 1;
}
```
6. generate support gRPC service file. protobuf generate service file from ```.proto```.
```shell
$ protoc --go_out=plugins=grpc:{output_path} {script_name}.proto
```
7. implement business logic.
8. start go script
```shell
$ go run {script_name}.go
```
