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
		HandlerType: (*transport.UserServer)(nil),
		Metadata:    "user.proto",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Create",
				Handler:    transport.MakeCreateHandler("/pb.User/Create"),
			},
			{
				MethodName: "Delete",
				Handler:    transport.MakeDeleteHandler("/pb.User/Delete"),
			},
		},
	}, transport.NewUserServer(svc.NewUserSvc()))
}
