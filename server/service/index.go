package service

import "whimsy/server/service/user"

type ServiceGroup struct {
	UserService user.UserService
}

var ServiceGroupApp = new(ServiceGroup)
