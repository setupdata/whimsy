package router

type RouterGroup struct {
	UserRouter  UserRouter
	VCodeRouter VCodeRouter
}

var RouterGroupApp = new(RouterGroup)
