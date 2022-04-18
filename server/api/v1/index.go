package v1

import (
	"whimsy/server/service"
)

type ApiGroup struct {
	UserApi  UserApi
	VCodeApi VCodeApi
}

var ApiGroupApp = new(ApiGroup)

var (
	userService   = service.ServiceGroupApp.UserService
	publicService = service.ServiceGroupApp.PublicService
)
