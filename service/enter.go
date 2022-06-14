package service

type ServiceGroup struct {
	UserService
	AuthorityService
	CasbinService
}

var ServiceGroupApp = new(ServiceGroup)
