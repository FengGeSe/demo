package user

import (
	"context"
	endpoint "demo/endpoint/user"
	pb "demo/pb/user"
	svc "demo/service/user"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// server
// UserServer is the server API for Cmd service.
type UserServer interface {
	Create(context.Context, *pb.CreateReq) (*pb.CreateResp, error)
	Delete(context.Context, *pb.DeleteReq) (*pb.DeleteResp, error)
}

type userServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
}

func (s userServer) Create(c context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	_, rep, err := s.create.ServeGRPC(c, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateResp), nil
}

func (s userServer) Delete(c context.Context, req *pb.DeleteReq) (*pb.DeleteResp, error) {
	_, rep, err := s.delete.ServeGRPC(c, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteResp), nil
}

func NewUserServer(svc svc.UserSvc, opts ...grpctransport.ServerOption) UserServer {
	return userServer{
		create: grpctransport.NewServer(
			endpoint.MakeCreateEndpoint(svc),
			decodeCreateRequest,
			encodeCreateResponse,
			opts...,
		),

		delete: grpctransport.NewServer(
			endpoint.MakeDeleteEndpoint(svc),
			decodeDeleteRequest,
			encodeDeleteResponse,
			opts...,
		),
	}
}
