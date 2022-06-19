package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
	NoticeRouter
	FileRouter
}

var RouterGroupApp = new(RouterGroup)
