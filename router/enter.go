package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
}

var RouterGroupApp = new(RouterGroup)
