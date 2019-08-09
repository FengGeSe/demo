package user

import (
	"context"
	model "demo/model/user"
)

// UserSvc is the server API for User service.
type UserSvc interface {
	Create(context.Context, *model.CreateReq) (*model.CreateResp, error)
	Delete(context.Context, *model.DeleteReq) (*model.DeleteResp, error)
}

// new service
func NewUserSvc() UserSvc {
	var svc = &userSvc{}
	{
		// middleware
	}

	return svc
}

// implement
type userSvc struct{}
