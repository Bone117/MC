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
		stageRouter.POST("deleteSign", stageApi.DeleteSign)
		stageRouter.POST("getSign", stageApi.GetSign)
		stageRouter.POST("getSignList", stageApi.GetSignList)

		stageRouter.POST("upload", stageApi.UploadFile)
		stageRouter.POST("deleteFile", stageApi.DeleteFile)
		stageRouter.GET("getFile", stageApi.GetFile)
		//stageRouter.GET("getVideoList", stageApi.GetVideoList)
	}
}
