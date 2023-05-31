package router

import (
	"server/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type StageRouter struct {
}

func (s *StageRouter) InitStageRouter(Router *gin.RouterGroup) {
	stageRouter := Router.Group("stage")
	stageApi := api.ApiGroupApp.StageApi
	{
		stageRouter.POST("sign", stageApi.Sign).Use(middleware.CheckStage())
		stageRouter.POST("updateSign", stageApi.UpdateSign).Use(middleware.CheckStage())
		stageRouter.POST("deleteSign", stageApi.DeleteSign).Use(middleware.CheckStage())
		stageRouter.GET("getSign", stageApi.GetSign)
		stageRouter.POST("getSignList", stageApi.GetSignList)
		stageRouter.GET("getSignEvaluate", stageApi.GetSignEvaluate)

		stageRouter.POST("upload", stageApi.UploadFile).Use(middleware.CheckStage())
		stageRouter.POST("deleteFile", stageApi.DeleteFile).Use(middleware.CheckStage())
		stageRouter.GET("getFile", stageApi.GetFile)
		stageRouter.GET("getUploadedFiles", stageApi.GetUploadedFiles)
		stageRouter.GET("getWorkFileType", stageApi.GetWorkFileType)
		stageRouter.GET("getMajor", stageApi.GetMajor)

		stageRouter.POST("/sign/:id/like", stageApi.Like)
	}
}
