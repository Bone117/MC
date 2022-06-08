package api

import "server/service"

type ApiGroup struct {
	BaseApi
	//UserApi
}

var (
	userService = service.ServiceGroupApp.UserService
)

var ApiGroupApp = new(ApiGroup)
