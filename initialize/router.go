package initialize

import (
	"server/global"
	"server/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	global.LOG.Info("use middleware logger")

	global.LOG.Info("register swagger handler")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := Router.Group("")
	systemRouter := router.RouterGroupApp
	systemRouter.InitBaseRouter(PublicGroup)

	global.LOG.Info("router register success")
	return Router
}
