package api

import (
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	Req "server/model/request"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PeriodApi struct {
}

//func (s *PeriodApi) CreatePeriod(ctx *gin.Context) {
//	periodR := model.CompetitionTime{}
//	_ = ctx.ShouldBindJSON(&periodR)
//	period := &model.CompetitionTime{Period: periodR.Period, StageID: periodR.StageID, StartTime: periodR.StartTime, EndTime: periodR.EndTime}
//	err := periodService.CreatePeriod(*period)
//	if err != nil {
//		global.LOG.Error("比赛届次创建失败!", zap.Error(err))
//		response.FailWithMessage("比赛届次创建失败", ctx)
//	} else {
//		response.OkWithDetailed(period, "比赛届次创建成功", ctx)
//	}
//}

func (s *PeriodApi) CreatePeriod(ctx *gin.Context) {
	periodR := Req.PeriodRequest{}
	_ = ctx.ShouldBindJSON(&periodR)

	stageTimeRanges := [][2]time.Time{
		{periodR.Stage1Starttime, periodR.Stage1Endtime},
		{periodR.Stage2Starttime, periodR.Stage2Endtime},
		{periodR.Stage3Starttime, periodR.Stage3Endtime},
		{periodR.Stage4Starttime, periodR.Stage4Endtime},
		{periodR.Stage5Starttime, periodR.Stage5Endtime},
	}

	// 遍历每个时间范围并创建 CompetitionTime 实例
	for i, timeRange := range stageTimeRanges {
		period := &model.CompetitionTime{
			Period:    periodR.Period,
			StageID:   uint(i) + 1,
			StartTime: timeRange[0],
			EndTime:   timeRange[1],
		}
		if err := periodService.CreatePeriod(*period); err != nil {
			global.LOG.Error("比赛阶段创建失败!", zap.Error(err))
			response.FailWithMessage("比赛创建失败", ctx)
			return
		}
	}
	response.OkWithMessage("比赛创建成功", ctx)
}

func (s *PeriodApi) UpdatePeriod(ctx *gin.Context) {
	periodR := model.CompetitionTime{}
	_ = ctx.ShouldBindJSON(&periodR)
	period := &model.CompetitionTime{Period: periodR.Period, StageID: periodR.StageID, StartTime: periodR.StartTime, EndTime: periodR.EndTime}
	err := periodService.CreatePeriod(*period)
	if err != nil {
		global.LOG.Error("比赛届次更新失败!", zap.Error(err))
		response.FailWithMessage("比赛届次更新失败", ctx)
	} else {
		response.OkWithMessage("比赛届次更新成功", ctx)
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
