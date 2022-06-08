package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("register", baseApi.Register)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
