package router

import (
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var http_router = map[string]http.Handler{}

// http register
func RegisterHttpRouter(path, handler) {
	http_router[path] = handler
}

var grpc_router = map[*grpc.ServiceDesc]interface{}{}

// grpc register
func RegisterGrpcRouter(sd *ServiceDesc, ss interface{}) {
	grpc_router[sd] = ss
}
