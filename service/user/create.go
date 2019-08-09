package user

import (
	"context"
	"fmt"

	model "demo/model/user"
)

func (s *userSvc) Create(context.Context, *model.CreateReq) (*model.CreateResp, error) {
	fmt.Println("Create")
	return model.NewCreateResp(), nil
}
