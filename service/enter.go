package service

type ServiceGroup struct {
	UserService
	AuthorityService
	NoticeService
	CasbinService
	StageService
	PeriodService
	ReviewService
}

var ServiceGroupApp = new(ServiceGroup)
