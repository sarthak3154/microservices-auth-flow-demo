package main

import (
  "context"
  "flag"
  "net/http"
  "log"

  "github.com/golang/glog"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

  gw "github.com/sarthak3154/microservices-auth-flow-demo/backend/protos"
)

var (
  // Command Line Options
  grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC Server Endpoint")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err1 := gw.RegisterPublicAccessServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
  if err1 != nil {
    return err1
  }

  err2 := gw.RegisterSecuredAccessServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
  if err2 != nil {
    return err2
  }

  log.Println("Starting GRPC Rest Gateway")
  return http.ListenAndServe(":8080", mux)
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }
}