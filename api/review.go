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
	reviewR := Req.ReviewRequest{}
	var err error
	_ = ctx.ShouldBindJSON(&reviewR)
	if err = utils.Verify(reviewR, utils.AssignVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	var reviewSigns []*model.ReviewSign

	period, err := stageService.GetJieCi()

	for _, signId := range reviewR.SignId {
		for _, userId := range reviewR.UserId {
			// 为每个评委和作品创建一个新的评估记录
			reviewSign := &model.ReviewSign{
				JieCiId:      period.Period,
				ReviewUserId: userId,
				SignId:       signId,
			}
			reviewSigns = append(reviewSigns, reviewSign)
		}
	}

	err = reviewService.CreateOrUpdateReviews(reviewSigns, reviewR.SignId, reviewR.UserId)

	if err != nil {
		global.LOG.Error("审核创建失败!", zap.Error(err))
		response.FailWithMessage("审核创建失败", ctx)
	} else {
		response.OkWithMessage("审核创建成功", ctx)
	}
}

func (r *ReviewApi) CreateEvaluate(ctx *gin.Context) {
	evaluateR := Req.EvaluateRequest{}
	var err error
	_ = ctx.ShouldBindJSON(&evaluateR)

	// 创建一个新的评估记录列表
	var evaluates []*model.Evaluate
	for _, signId := range evaluateR.SignId {
		for _, userId := range evaluateR.UserId {
			// 为每个评委和作品创建一个新的评估记录
			evaluate := &model.Evaluate{
				JieCiId:        evaluateR.JieCiId,
				EvaluateUserId: userId,
				SignId:         signId,
			}
			evaluates = append(evaluates, evaluate)
		}
	}

	// 调用服务层的CreateEvaluates方法批量创建或更新评估记录
	err = reviewService.CreateOrUpdateEvaluates(evaluates, evaluateR.SignId, evaluateR.UserId)

	if err != nil {
		global.LOG.Error("评审创建失败!", zap.Error(err))
		response.FailWithMessage("评审创建失败", ctx)
	} else {
		response.OkWithMessage("评审创建成功", ctx)
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

func (r *ReviewApi) GetEvaluateList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	userID := utils.GetUserID(ctx)
	if list, total, err := reviewService.GetEvaluateList(pageInfo, userID); err != nil {
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

func (r *ReviewApi) UpdateEvaluate(ctx *gin.Context) {
	evaluateR := Req.UpdateEvaluateRequest{}
	_ = ctx.ShouldBindJSON(&evaluateR)
	evaluate := &model.Evaluate{SignId: evaluateR.SignId, EvaluateUserId: utils.GetUserID(ctx), Score: evaluateR.Score, Comments: evaluateR.Comments}
	err := reviewService.UpdateEvaluate(*evaluate)
	if err != nil {
		global.LOG.Error("评分更新失败!", zap.Error(err))
		response.FailWithMessage("评分失败", ctx)
	} else {
		response.OkWithMessage("评分成功", ctx)
	}
}

func (r *ReviewApi) CreateReport(ctx *gin.Context) {
	reportR := Req.ReportRequest{}
	var err error
	_ = ctx.ShouldBindJSON(&reportR)

	report := &model.Report{ReportUserId: utils.GetUserID(ctx), JieCiId: reportR.JieCiId, SignId: reportR.SignId, Content: reportR.Content}
	err = reviewService.CreateOrUpdateReport(*report)

	if err != nil {
		global.LOG.Error("举报创建失败!", zap.Error(err))
		response.FailWithMessage("举报提交失败", ctx)
	} else {
		response.OkWithMessage("举报提交成功", ctx)
	}
}

func (r *ReviewApi) GetReportList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if list, total, err := reviewService.GetReportList(pageInfo); err != nil {
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
