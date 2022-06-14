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
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup) // 注册用户路由

	}

	global.LOG.Info("router register success")
	return Router
}
