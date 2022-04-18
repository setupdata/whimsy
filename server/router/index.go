package router

import "whimsy/server/router/user"

type RouterGroup struct {
	//ApiRouter
	//JwtRouter
	//SysRouter
	//BaseRouter
	//InitRouter
	//MenuRouter
	UserRouter user.UserRouter
	//CasbinRouter
	//AutoCodeRouter
	//AuthorityRouter
	//DictionaryRouter
	//OperationRecordRouter
	//DictionaryDetailRouter
}

var RouterGroupApp = new(RouterGroup)
