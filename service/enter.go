package service

type ServiceGroup struct {
	UserService
	AuthorityService
	NoticeService
	CasbinService
	StageService
}

var ServiceGroupApp = new(ServiceGroup)
