package api

import "server/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	NoticeApi
	StageApi
}

var (
	userService      = service.ServiceGroupApp.UserService
	authorityService = service.ServiceGroupApp.AuthorityService
	casbinService    = service.ServiceGroupApp.CasbinService
	noticeService    = service.ServiceGroupApp.NoticeService
	StageService     = service.ServiceGroupApp.StageService
)

var ApiGroupApp = new(ApiGroup)
