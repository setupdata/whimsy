package request

type RequestGroup struct {
	User UserReq
}

var RequestGroupApp = new(RequestGroup)
