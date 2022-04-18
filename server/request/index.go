package request

import (
	"whimsy/server/request/user"
)

type RequestGroup struct {
	User user.UserReq
}

var RequestGroupApp = new(RequestGroup)
