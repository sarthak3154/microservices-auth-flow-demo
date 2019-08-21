#!/usr/bin/env bash

if [[ "$#" -ne 2 ]] || [[ "$1" != "-f" ]]; then
    echo "USAGE: ./genproto.sh -f proto-file-name"
    exit 1
fi

FILE_NAME="$2"

if [[ ${FILE_NAME: -6} != ".proto" ]]; then
    echo "Invalid file type. Please input a '.proto' file name"
    exit 1
fi

# Generate the gRPC stub for Go, ending with .pb.go
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. protos/${FILE_NAME}
echo "Generated gRPC stub for ${FILE_NAME}"

# Generate the gRPC reverse proxy in Go, ending with .pb.gw.go
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. protos/${FILE_NAME}
echo "Generated gRPC reverse proxy for ${FILE_NAME}"