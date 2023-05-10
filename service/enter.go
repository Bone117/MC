package service

type ServiceGroup struct {
	UserService
	AuthorityService
	NoticeService
	CasbinService
	StageService
	PeriodService
	ReviewService
	PortfolioService
}

var ServiceGroupApp = new(ServiceGroup)
