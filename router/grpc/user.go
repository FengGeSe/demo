package grpc

import (
	server "demo/server/grpc"
	svc "demo/service/user"
	transport "demo/transport/grpc/user"

	"google.golang.org/grpc"
)

func init() {
	registerUserService()
}

func registerUserService() {
	server.RegisterService(&grpc.ServiceDesc{
		ServiceName: "pb.User",
		HandlerType: (*svc.UserSvc)(nil),
		Metadata:    "user.proto",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Create",
				Handler:    transport.UserCreateHandler,
			},
			{
				MethodName: "Delete",
				Handler:    transport.UserDeleteHandler,
			},
		},
	}, svc.NewUserSvc())
}
