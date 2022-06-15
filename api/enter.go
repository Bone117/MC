package api

import "server/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
}

var (
	userService      = service.ServiceGroupApp.UserService
	authorityService = service.ServiceGroupApp.AuthorityService
	casbinService    = service.ServiceGroupApp.CasbinService
)

var ApiGroupApp = new(ApiGroup)
