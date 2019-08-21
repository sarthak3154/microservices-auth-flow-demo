#!/usr/bin/env bash

# Generate the gRPC stub for Go, ending with .pb.go
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. protos/message.proto

# Generate the gRPC reverse proxy in Go, ending with .pb.gw.go
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. protos/message.proto