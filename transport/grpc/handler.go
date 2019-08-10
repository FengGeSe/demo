package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type MethodHandler func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error)
