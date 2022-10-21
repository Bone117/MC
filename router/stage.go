package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type StageRouter struct {
}

func (s *StageRouter) InitStageRouter(Router *gin.RouterGroup) {
	stageRouter := Router.Group("stage")
	stageApi := api.ApiGroupApp.StageApi
	{
		stageRouter.POST("sign", stageApi.Sign)
		stageRouter.POST("updateSign", stageApi.UpdateSign)
		stageRouter.POST("deleteSign", stageApi.Sign)
		stageRouter.POST("getSign", stageApi.Sign)
		stageRouter.POST("getSignList", stageApi.Sign)

		//fileRouter.POST("upload", fileApi.UploadFile)
	}
}
