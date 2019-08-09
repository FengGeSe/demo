package grpc

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var opts = []grpc.ServerOption{
	grpc_middleware.WithUnaryServerChain(
		RecoveryInterceptor,
	),
}

var grpcServer = grpc.NewServer(opts...)

func RegisterService(sd *grpc.ServiceDesc, svc interface{}) {
	grpcServer.RegisterService(sd, svc)
}

func Run(addr string, errc chan error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}

	errc <- grpcServer.Serve(lis)
}
