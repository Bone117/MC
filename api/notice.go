package api

import (
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	Req "server/model/request"
	Res "server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NoticeApi struct{}

func (n *NoticeApi) CreateNotice(ctx *gin.Context) {
	noticeR := model.Notice{}
	_ = ctx.ShouldBindJSON(&noticeR)
	if err := utils.Verify(noticeR, utils.NoticeVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	notice := &model.Notice{Title: noticeR.Title, Desc: noticeR.Desc, Content: noticeR.Content}
	id, err := noticeService.CreateNotice(*notice)
	if err != nil {
		global.LOG.Error("公告创建失败!", zap.Error(err))
		response.FailWithDetailed(Res.NoticeResponse{NoticeId: id}, "公告创建失败", ctx)
	} else {
		response.OkWithDetailed(Res.NoticeResponse{NoticeId: id}, "公告创建成功", ctx)
	}
}

func (n *NoticeApi) DeleteNotice(ctx *gin.Context) {
	notice := Req.NoticeRequest{}
	_ = ctx.ShouldBindJSON(&notice)
	if err := utils.Verify(notice, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := noticeService.DeleteNotice(notice.NoticeId); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}

}

func (n *NoticeApi) UpdateNotice(ctx *gin.Context) {
	noticeR := model.Notice{}
	_ = ctx.ShouldBindJSON(&noticeR)
	if err := utils.Verify(noticeR, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	notice := &model.Notice{MODEL: global.MODEL{ID: noticeR.ID}, Title: noticeR.Title, Desc: noticeR.Desc, Content: noticeR.Content}
	err := noticeService.UpdateNotice(*notice)
	if err != nil {
		global.LOG.Error("公告更新失败!", zap.Error(err))
		response.FailWithDetailed(Res.NoticeResponse{NoticeId: notice.ID}, "公告更新失败", ctx)
	} else {
		response.OkWithDetailed(Res.NoticeResponse{NoticeId: notice.ID}, "公告更新成功", ctx)
	}
}

func (n *NoticeApi) GetNotice(ctx *gin.Context) {
	noticeR := Req.NoticeRequest{}
	_ = ctx.ShouldBindQuery(&noticeR)
	if err := utils.Verify(noticeR, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	notice, err := noticeService.GetNotice(noticeR.NoticeId)
	if err != nil {
		global.LOG.Error("公告获取失败!", zap.Error(err))
		response.FailWithDetailed(notice, "公告获取失败", ctx)
	} else {
		response.OkWithDetailed(notice, "公告获取成功", ctx)
	}
}

func (n *NoticeApi) GetNoticeList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := noticeService.GetNoticeList(pageInfo); err != nil {
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
