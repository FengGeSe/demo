package user

import (
	"context"
	"fmt"

	model "demo/model/user"
)

func (s *userSvc) Create(ctx context.Context, req *model.CreateReq) (*model.CreateResp, error) {
	fmt.Println("Create")
	resp := model.NewCreateResp()
	resp.Code = 200
	resp.Msg = "success"
	resp.Data.Id = "12"
	resp.Data.Name = req.Name
	resp.Data.Age = req.Age
	return resp, nil
}
