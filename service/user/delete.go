package user

import (
	"context"
	"fmt"

	model "demo/model/user"
)

func (s *userSvc) Delete(context.Context, *model.DeleteReq) (*model.DeleteResp, error) {
	fmt.Println("Delete")
	return model.NewDeleteResp(), nil
}
