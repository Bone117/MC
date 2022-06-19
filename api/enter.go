package api

import "server/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	NoticeApi
	FileApi
}

var (
	userService      = service.ServiceGroupApp.UserService
	authorityService = service.ServiceGroupApp.AuthorityService
	casbinService    = service.ServiceGroupApp.CasbinService
	noticeService    = service.ServiceGroupApp.NoticeService
)

var ApiGroupApp = new(ApiGroup)
