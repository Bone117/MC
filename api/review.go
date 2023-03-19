package api

import (
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	Req "server/model/request"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReviewApi struct {
}

func (r *ReviewApi) CreateReview(ctx *gin.Context) {
	//TODO 需要再修改
	reviewR := Req.ReviewRequest{}
	var err error
	_ = ctx.ShouldBindJSON(&reviewR)
	//for _, signId := range reviewR.SignId {
	//	for _, userId := range reviewR.UserId {
	//		review := &model.ReviewSign{ReviewUserId: userId, SignId: signId, JieCiId: reviewR.JieCiId}
	//		err = reviewService.CreateReview(*review)
	//	}
	//}
	review := &model.ReviewSign{JieCiId: reviewR.JieCiId}

	for _, signId := range reviewR.SignId {
		for _, userId := range reviewR.UserId {
			// set the ReviewUserId and SignId fields for each iteration
			review.ReviewUserId = userId
			review.SignId = signId
			err = reviewService.CreateReview(*review)
		}
	}

	if err != nil {
		global.LOG.Error("审核创建失败!", zap.Error(err))
		response.FailWithMessage("审核创建失败", ctx)
	} else {
		response.OkWithMessage("审核创建成功", ctx)
	}
}

func (r *ReviewApi) DeleteReview(ctx *gin.Context) {
	review := Req.ReviewRequest{}
	_ = ctx.ShouldBindJSON(&review)
	//if err := utils.Verify(reviewR, utils.NoticeVerify); err != nil {
	//	response.FailWithMessage(err.Error(), ctx)
	//}
	if err := reviewService.DeleteReview(review.ReviewId); err != nil {
		global.LOG.Error("审核删除失败!", zap.Error(err))
		response.FailWithMessage("审核删除失败", ctx)
	} else {
		response.OkWithMessage("审核删除成功", ctx)
	}
}

func (r *ReviewApi) UpdateReview(ctx *gin.Context) {
	reviewR := model.ReviewSign{}
	_ = ctx.ShouldBindJSON(&reviewR)
	//if err := utils.Verify(reviewR, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), ctx)
	//}
	review := &model.ReviewSign{MODEL: global.MODEL{ID: reviewR.ID}, ReviewUserId: reviewR.ReviewUserId, SignId: reviewR.SignId, JieCiId: reviewR.JieCiId}
	err := reviewService.UpdateReview(*review)
	if err != nil {
		global.LOG.Error("审核更新失败!", zap.Error(err))
		response.FailWithMessage("审核更新失败", ctx)
	} else {
		response.OkWithMessage("审核更新成功", ctx)
	}
}

func (r *ReviewApi) GetReviewList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := reviewService.GetReviewList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("暂未分配", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}

}
