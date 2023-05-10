package api

import (
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PortfolioApi struct {
}

func (p *PortfolioApi) UploadFile(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}
	userid := utils.GetUserID(ctx)
	portfolioTitle := ctx.PostForm("title")
	filePath, uploadErr := utils.UploadFile(header, userid, 0) // 文件上传后拿到文件路径
	if uploadErr != nil {
		panic(uploadErr)
	}

	// 生成缩略图并更新Sign表
	var coverUrl string
	coverUrl, err = utils.GenerateThumbnail(filePath)
	if err != nil {
		global.LOG.Error("生成缩略图失败!", zap.Error(err))
		return
	}
	f1 := model.Portfolio{
		Url:      filePath,
		Title:    portfolioTitle,
		FileName: header.Filename,
		CoverUrl: coverUrl,
	}

	if file, err := portfolioService.Upload(f1); err != nil {
		response.FailWithMessage("上传失败", ctx)
		return
	} else {
		response.OkWithDetailed(file, "上传成功", ctx)
	}
}

func (p *PortfolioApi) GetPastWork(ctx *gin.Context) {
	var reqId request.GetById
	_ = ctx.ShouldBindQuery(&reqId)
	if sign, err := portfolioService.GetPastWork(reqId.ID); err != nil {
		global.LOG.Error("往届作品信息获取失败!", zap.Error(err))
		response.FailWithDetailed(sign, "作品信息获取失败", ctx)
	} else {
		response.OkWithDetailed(sign, "作品信息获取成功", ctx)
	}
}

func (p *PortfolioApi) GetPastWorkList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)

	if list, total, err := portfolioService.GetPortfolioList(pageInfo); err != nil {
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
