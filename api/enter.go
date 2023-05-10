package api

import "server/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	NoticeApi
	StageApi
	PeriodApi
	ReviewApi
	PortfolioApi
}

var (
	userService      = service.ServiceGroupApp.UserService
	authorityService = service.ServiceGroupApp.AuthorityService
	casbinService    = service.ServiceGroupApp.CasbinService
	noticeService    = service.ServiceGroupApp.NoticeService
	stageService     = service.ServiceGroupApp.StageService
	portfolioService = service.ServiceGroupApp.PortfolioService
	periodService    = service.ServiceGroupApp.PeriodService
	reviewService    = service.ServiceGroupApp.ReviewService
)

var ApiGroupApp = new(ApiGroup)
