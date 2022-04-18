package service

type ServiceGroup struct {
	UserService
	PublicService
}

var ServiceGroupApp = new(ServiceGroup)
