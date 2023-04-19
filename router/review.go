package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type ReviewRouter struct{}

func (p *PeriodRouter) InitReviewRouter(Router *gin.RouterGroup) {
	reviewRouter := Router.Group("review")
	reviewApi := api.ApiGroupApp.ReviewApi
	{
		reviewRouter.POST("createReview", reviewApi.CreateReview)
		reviewRouter.POST("deleteReview", reviewApi.DeleteReview)
		reviewRouter.POST("updateReview", reviewApi.UpdateReview)
		reviewRouter.POST("getReviewList", reviewApi.GetReviewList)
		reviewRouter.POST("getEvaluateList", reviewApi.GetEvaluateList)

		reviewRouter.POST("createEvaluate", reviewApi.CreateEvaluate)
		reviewRouter.POST("updateEvaluate", reviewApi.UpdateEvaluate)

		reviewRouter.POST("createReport", reviewApi.CreateReport)
		reviewRouter.POST("getReportList", reviewApi.GetReportList)
	}
}
