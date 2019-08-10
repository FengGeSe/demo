package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	errors "demo/errors"
	model "demo/model/user"
	svc "demo/service/user"
)

// make endpoint             service -> endpoint
func MakeDeleteEndpoint(svc svc.UserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*model.DeleteReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Delete(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
