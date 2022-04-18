package v1

import (
	"whimsy/server/api/v1/user"
	"whimsy/server/service"
)

type ApiGroup struct {
	UserApi user.Api
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = service.ServiceGroupApp.UserService
)
