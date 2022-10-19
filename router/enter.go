package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
	NoticeRouter
	StageRouter
}

var RouterGroupApp = new(RouterGroup)
