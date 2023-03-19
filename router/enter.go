package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
	NoticeRouter
	StageRouter
	PeriodRouter
	ReviewRouter
}

var RouterGroupApp = new(RouterGroup)
