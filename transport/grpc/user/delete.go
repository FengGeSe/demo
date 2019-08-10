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
func decodeDeleteRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.DeleteReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	request := &model.DeleteReq{
		Name: req.Name,
		Id:   req.Id,
	}
	return request, nil
}

// 2. encode response           model -> pb
func encodeDeleteResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.DeleteResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response出错！")
	}
	r := &pb.DeleteResp{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: &pb.DeleteRespData{
			Result: resp.Data.Result,
		},
	}
	return r, nil
}

func MakeDeleteHandler(fullMethod string) func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	return func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		in := new(pb.DeleteReq)
		if err := dec(in); err != nil {
			return nil, err
		}
		if interceptor == nil {
			return srv.(UserServer).Delete(ctx, in)
		}
		info := &grpc.UnaryServerInfo{
			Server:     srv,
			FullMethod: fullMethod,
		}
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(UserServer).Delete(ctx, req.(*pb.DeleteReq))
		}
		return interceptor(ctx, in, info, handler)
	}
}
