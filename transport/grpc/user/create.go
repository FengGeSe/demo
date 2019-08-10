package user

import (
	"context"
	"fmt"

	model "demo/model/user"
	pb "demo/pb/user"

	"google.golang.org/grpc"
)

// Server
// 1. decode request          pb -> model
func decodeCreateRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.CreateReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	request := &model.CreateReq{
		Name: req.Name,
		Age:  req.Age,
	}
	return request, nil
}

// 2. encode response           model -> pb
func encodeCreateResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.CreateResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", response)
	}
	r := &pb.CreateResp{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: &pb.CreateRespData{
			Id:   resp.Data.Id,
			Age:  resp.Data.Age,
			Name: resp.Data.Name,
		},
	}
	return r, nil
}

func MakeCreateHandler(fullMethod string) func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	return func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		in := new(pb.CreateReq)
		if err := dec(in); err != nil {
			return nil, err
		}
		if interceptor == nil {
			return srv.(UserServer).Create(ctx, in)
		}
		info := &grpc.UnaryServerInfo{
			Server:     srv,
			FullMethod: fullMethod,
		}
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(UserServer).Create(ctx, req.(*pb.CreateReq))
		}
		return interceptor(ctx, in, info, handler)
	}
}
