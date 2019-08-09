package user

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	model "demo/model/user"
	service "demo/service/user"
	transport "demo/transport/http"

	endpoint "demo/endpoint/user"
)

// Server
// 1. decode request      http.request -> model.request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if err := transport.FormCheckAccess(r); err != nil {
		return nil, err
	}
	r.ParseForm()
	req := &model.CreateReq{}
	err := transport.ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return req, nil
}

// 2. encode response      model.response -> http.response
func encodeCreateResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

// make handler
func MakeCreateHandler(svc service.UserSvc) http.Handler {
	handler := httptransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}
