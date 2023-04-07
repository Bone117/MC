package router

import (
	"server/api"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type NoticeRouter struct{}

func (n *NoticeRouter) InitNoticeRouter(Router *gin.RouterGroup) {
	noticeRouter := Router.Group("notice").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	noticeRouterWithoutLoginL := Router.Group("notice")
	noticeApi := api.ApiGroupApp.NoticeApi
	{
		noticeRouter.POST("createNotice", noticeApi.CreateNotice)
		noticeRouter.POST("deleteNotice", noticeApi.DeleteNotice)
		noticeRouter.POST("updateNotice", noticeApi.UpdateNotice)
	}
	{
		noticeRouterWithoutLoginL.GET("getNotice", noticeApi.GetNotice)
		noticeRouterWithoutLoginL.POST("getNoticeList", noticeApi.GetNoticeList)
	}
}
