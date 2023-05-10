package initialize

import (
	"server/global"
	"server/middleware"
	"server/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := router.RouterGroupApp
	global.LOG.Info("use middleware logger")

	global.LOG.Info("register swagger handler")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitNoticeRouter(PublicGroup)
		systemRouter.InitPortfolioRouter(PublicGroup) // 注册优秀作品路由
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)      // 注册用户路由
		systemRouter.InitAuthorityRouter(PrivateGroup) // 注册角色路由
		systemRouter.InitStageRouter(PrivateGroup)     // 注册文件路由
		systemRouter.InitPeriodRouter(PrivateGroup)    // 注册比赛届次路由
		systemRouter.InitReviewRouter(PrivateGroup)    // 注册审核路由
	}

	global.LOG.Info("router register success")
	return Router
}
