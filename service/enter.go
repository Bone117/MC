package service

type ServiceGroup struct {
	UserService
	AuthorityService
	NoticeService
	CasbinService
}

var ServiceGroupApp = new(ServiceGroup)
