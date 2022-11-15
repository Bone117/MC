package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
	NoticeRouter
	StageRouter
	PeriodRouter
}

var RouterGroupApp = new(RouterGroup)
