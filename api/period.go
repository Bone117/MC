package api

import (
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PeriodApi struct {
}

func (s *PeriodApi) CreatePeriod(ctx *gin.Context) {
	periodR := model.CompetitionTime{}
	_ = ctx.ShouldBindJSON(&periodR)
	period := &model.CompetitionTime{Period: periodR.Period, StageID: periodR.StageID, StartTime: periodR.StartTime, EndTime: periodR.EndTime}
	err := periodService.CreatePeriod(*period)
	if err != nil {
		global.LOG.Error("比赛届次创建失败!", zap.Error(err))
		response.FailWithMessage("比赛届次创建失败", ctx)
	} else {
		response.OkWithDetailed(period, "比赛届次创建成功", ctx)
	}
}

func (s *PeriodApi) UpdatePeriod(ctx *gin.Context) {
	periodR := model.Period{}
	_ = ctx.ShouldBindJSON(&periodR)
	period := &model.Period{MODEL: global.MODEL{ID: periodR.ID}, JieCi: periodR.JieCi}
	err := periodService.UpdatePeriod(*period)
	if err != nil {
		global.LOG.Error("比赛届次更新失败!", zap.Error(err))
		response.FailWithMessage("比赛届次更新失败", ctx)
	} else {
		response.FailWithMessage("比赛届次更新成功", ctx)
	}
}

func (s *PeriodApi) DeletePeriod(ctx *gin.Context) {
	var reqId request.GetById
	_ = ctx.ShouldBindJSON(&reqId)
	if err := periodService.DeletePeriod(reqId.ID); err != nil {
		global.LOG.Error("届次删除失败!", zap.Error(err))
		response.FailWithMessage("届次删除失败", ctx)
	} else {
		response.OkWithMessage("届次删除成功", ctx)
	}
}

func (s *PeriodApi) GetPeriod(ctx *gin.Context) {
	var reqId request.GetById
	_ = ctx.ShouldBindJSON(&reqId)
	if period, err := periodService.GetPeriod(reqId.ID); err != nil {
		global.LOG.Error("比赛届次获取失败!", zap.Error(err))
		response.FailWithDetailed(period, "比赛届次获取失败", ctx)
	} else {
		response.OkWithDetailed(period, "比赛届次获取成功", ctx)
	}
}

//func (s *PeriodApi) GetCurrentPeriod(ctx *gin.Context) {
//	var reqId request.GetById
//	_ = ctx.ShouldBindJSON(&reqId)
//	if period, err := periodService.GetPeriod(reqId.ID); err != nil {
//		global.LOG.Error("比赛届次获取失败!", zap.Error(err))
//		response.FailWithDetailed(period, "比赛届次获取失败", ctx)
//	} else {
//		response.OkWithDetailed(period, "比赛届次获取成功", ctx)
//	}
//}

func (s *PeriodApi) GetPeriodList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if list, total, err := periodService.GetPeriodList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

func (s *PeriodApi) CreateCpTime(ctx *gin.Context) {

}

func (s *PeriodApi) UpdateCpTime(ctx *gin.Context) {

}

func (s *PeriodApi) DeleteCpTime(ctx *gin.Context) {

}

func (s *PeriodApi) GetCpTimeList(ctx *gin.Context) {

}
