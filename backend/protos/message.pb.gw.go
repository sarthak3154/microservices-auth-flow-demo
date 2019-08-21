// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: protos/message.proto

/*
Package message is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package message

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_PublicAccessService_GetMessage_0(ctx context.Context, marshaler runtime.Marshaler, client PublicAccessServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetMessage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_SecuredAccessService_GetMessage_0(ctx context.Context, marshaler runtime.Marshaler, client SecuredAccessServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetMessage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterPublicAccessServiceHandlerFromEndpoint is same as RegisterPublicAccessServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPublicAccessServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterPublicAccessServiceHandler(ctx, mux, conn)
}

// RegisterPublicAccessServiceHandler registers the http handlers for service PublicAccessService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPublicAccessServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPublicAccessServiceHandlerClient(ctx, mux, NewPublicAccessServiceClient(conn))
}

// RegisterPublicAccessServiceHandlerClient registers the http handlers for service PublicAccessService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PublicAccessServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PublicAccessServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PublicAccessServiceClient" to call the correct interceptors.
func RegisterPublicAccessServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PublicAccessServiceClient) error {

	mux.Handle("GET", pattern_PublicAccessService_GetMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PublicAccessService_GetMessage_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_PublicAccessService_GetMessage_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_PublicAccessService_GetMessage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"public"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_PublicAccessService_GetMessage_0 = runtime.ForwardResponseMessage
)

// RegisterSecuredAccessServiceHandlerFromEndpoint is same as RegisterSecuredAccessServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSecuredAccessServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSecuredAccessServiceHandler(ctx, mux, conn)
}

// RegisterSecuredAccessServiceHandler registers the http handlers for service SecuredAccessService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSecuredAccessServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSecuredAccessServiceHandlerClient(ctx, mux, NewSecuredAccessServiceClient(conn))
}

// RegisterSecuredAccessServiceHandlerClient registers the http handlers for service SecuredAccessService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SecuredAccessServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SecuredAccessServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SecuredAccessServiceClient" to call the correct interceptors.
func RegisterSecuredAccessServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SecuredAccessServiceClient) error {

	mux.Handle("GET", pattern_SecuredAccessService_GetMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SecuredAccessService_GetMessage_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SecuredAccessService_GetMessage_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SecuredAccessService_GetMessage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"secured"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_SecuredAccessService_GetMessage_0 = runtime.ForwardResponseMessage
)
