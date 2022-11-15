package api

import "server/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	NoticeApi
	StageApi
	PeriodApi
}

var (
	userService      = service.ServiceGroupApp.UserService
	authorityService = service.ServiceGroupApp.AuthorityService
	casbinService    = service.ServiceGroupApp.CasbinService
	noticeService    = service.ServiceGroupApp.NoticeService
	stageService     = service.ServiceGroupApp.StageService
	periodService    = service.ServiceGroupApp.PeriodService
)

var ApiGroupApp = new(ApiGroup)
