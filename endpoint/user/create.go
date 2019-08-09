package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	errors "demo/errors"
	model "demo/model/user"
	service "demo/service/user"
)

// make endpoint             service -> endpoint
func MakeCreateEndpoint(svc service.UserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*model.CreateReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
