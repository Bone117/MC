package service

type ServiceGroup struct {
	UserService
	AuthorityService
	NoticeService
	CasbinService
	StageService
	PeriodService
}

var ServiceGroupApp = new(ServiceGroup)
