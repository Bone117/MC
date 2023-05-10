package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	AuthorityRouter
	NoticeRouter
	StageRouter
	PeriodRouter
	ReviewRouter
	PortfolioRouter
}

var RouterGroupApp = new(RouterGroup)
