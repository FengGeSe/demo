package http

import (
	server "demo/server/http"
	svc "demo/service/user"
	transport "demo/transport/http/user"
)

func init() {
	registerUserHandler()
}

func registerUserHandler() {
	server.RegisterHandler("/user/create", transport.MakeCreateHandler(svc.NewUserSvc()))
	server.RegisterHandler("/user/delete", transport.MakeDeleteHandler(svc.NewUserSvc()))
}
