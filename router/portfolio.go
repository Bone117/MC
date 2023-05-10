package router

import (
	"server/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type PortfolioRouter struct {
}

func (p *PortfolioRouter) InitPortfolioRouter(Router *gin.RouterGroup) {
	portfolioRouter := Router.Group("portfolio")
	portfolioApi := api.ApiGroupApp.PortfolioApi
	{
		portfolioRouter.GET("getPastWork", portfolioApi.GetPastWork)
		portfolioRouter.POST("getPastWorkList", portfolioApi.GetPastWorkList)
		portfolioRouter.POST("upload", portfolioApi.UploadFile).Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	}
}
