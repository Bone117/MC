package api

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/model/common/response"
	Req "server/model/request"
	Res "server/model/response"
	"server/utils"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StageApi struct {
}

func (s *StageApi) Sign(ctx *gin.Context) {
	var signReq Req.SignRequest
	_ = ctx.ShouldBindJSON(&signReq)
	if err := utils.Verify(signReq, utils.SignVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	period, err := stageService.GetJieCi()
	if err != nil {
		global.LOG.Error("获取届次失败!", zap.Error(err))
		response.FailWithMessage("比赛暂未开始", ctx)
		return
	}
	sign := &model.Sign{
		UserId:         utils.GetUserID(ctx),
		WorkName:       signReq.WorkName,
		WorkFileTypeId: signReq.WorkFileTypeId,
		Author:         signReq.Author,
		Username:       signReq.Username,
		WorkSoftware:   signReq.WorkSoftware,
		OtherAuthor:    signReq.OtherAuthor,
		WorkAdviser:    signReq.WorkAdviser,
		WorkDesc:       signReq.WorkDesc,
		JieCiId:        period.Period,
	}
	gra := &model.Grade{
		UserId:    utils.GetUserID(ctx),
		MajorId:   signReq.MajorId,
		GradeName: signReq.GradeName,
	}
	if err := stageService.Sign(*sign, *gra); err != nil {
		global.LOG.Error("报名失败!", zap.Error(err))
		response.FailWithMessage("报名失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("报名成功", ctx)
	}

}

func (s *StageApi) UpdateSign(ctx *gin.Context) {
	signR := model.Sign{}
	_ = ctx.ShouldBindJSON(&signR)
	if err := utils.Verify(signR, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	sign := &model.Sign{MODEL: global.MODEL{ID: signR.ID}, WorkName: signR.WorkName, WorkFileTypeId: signR.WorkFileTypeId,
		OtherAuthor: signR.OtherAuthor, WorkAdviser: signR.WorkAdviser, WorkSoftware: signR.WorkSoftware,
		WorkDesc: signR.WorkDesc, Status: signR.Status, RejReason: signR.RejReason}
	if err := stageService.UpdateSign(*sign); err != nil {
		global.LOG.Error("报名更新失败!", zap.Error(err))
		response.FailWithMessage("报名更新失败", ctx)
	} else {
		response.OkWithMessage("报名更新成功", ctx)
	}
}

func (s *StageApi) DeleteSign(ctx *gin.Context) {
	var reqId request.GetById
	_ = ctx.ShouldBindJSON(&reqId)
	if err := stageService.DeleteSign(reqId.ID); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

func (s *StageApi) GetSign(ctx *gin.Context) {
	var reqId request.GetById
	_ = ctx.ShouldBindQuery(&reqId)
	if sign, err := stageService.GetSign(reqId.ID); err != nil {
		global.LOG.Error("报名信息获取失败!", zap.Error(err))
		response.FailWithDetailed(sign, "报名信息获取失败", ctx)
	} else {
		response.OkWithDetailed(sign, "报名信息获取成功", ctx)
	}
}

func (s *StageApi) GetSignList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := stageService.GetSignList(pageInfo); err != nil {
		global.LOG.Error("获取报名列表失败!", zap.Error(err))
		response.FailWithMessage("获取报名列表失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取报名列表成功", ctx)
	}
}

// UploadFile
// @Tags     stage
// @Summary  上传文件
// @Produce   multipart/form-data
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /stage/upload [post]
func (s *StageApi) UploadFile(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}
	var f1 model.File
	userid := utils.GetUserID(ctx)
	period, err := stageService.GetJieCi()
	if err != nil {
		global.LOG.Error("获取届次信息失败!", zap.Error(err))
		return
	}
	fileTypeIDStr := ctx.PostForm("fileTypeID")
	signIDStr := ctx.PostForm("signId")
	if fileTypeIDStr != "" && signIDStr != "" {
		fileTypeID, _ := strconv.Atoi(fileTypeIDStr)
		signID, _ := strconv.Atoi(signIDStr)
		keyWords := map[string]interface{}{
			"FileTypeID": uint(fileTypeID),
			"SignId":     uint(signID),
		}
		file, err := stageService.GetFile(keyWords)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithDetailed(file, "文件已存在", ctx)
			return
		}

		filePath, uploadErr := utils.UploadFile(header, userid, period.Period) // 文件上传后拿到文件路径
		if uploadErr != nil {
			panic(uploadErr)
		}

		// 生成缩略图并更新Sign表
		var coverUrl string
		if fileTypeID == 3 {
			coverUrl, err = utils.GenerateThumbnail(filePath)
			if err != nil {
				global.LOG.Error("生成缩略图失败!", zap.Error(err))
				return
			}
			err = stageService.UpdateSignCoverUrl(uint(signID), coverUrl)
			if err != nil {
				global.LOG.Error("Sign表生成缩略图失败!", zap.Error(err))
				response.FailWithMessage("Sign表生成缩略图失败", ctx)
				return
			}
		}

		f1 = model.File{
			UserId:     utils.GetUserID(ctx),
			Url:        filePath,
			FileName:   header.Filename,
			FileTypeID: uint(fileTypeID),
			SignId:     uint(signID),
		}
	} else {
		filePath, uploadErr := utils.UploadFile(header, userid, period.Period) // 文件上传后拿到文件路径
		if uploadErr != nil {
			panic(err)
		}
		f1 = model.File{
			UserId:   utils.GetUserID(ctx),
			Url:      filePath,
			FileName: header.Filename,
		}
	}

	if file, err := stageService.Upload(f1); err != nil {
		response.FailWithMessage("上传失败", ctx)
		return
	} else {
		response.OkWithDetailed(file, "上传成功", ctx)
	}
}

// GetFile
// @Tags     stage
// @Summary  获取文件
// @Param    data  param      Req.GetFileRequest                          true  "文件ID"
// @Success  200   {object}  response.Response{data=file.File,msg=string}  "返回文件信息"
// @Router   /stage/getFile [get]
func (s *StageApi) GetFile(ctx *gin.Context) {
	fileR := Req.GetFileRequest{}
	_ = ctx.ShouldBindQuery(&fileR)
	keyWords := map[string]interface{}{
		"id": fileR.FileId,
	}
	if file, err := stageService.GetFile(keyWords); err != nil {
		global.LOG.Error("文件信息获取失败!", zap.Error(err))
		response.FailWithDetailed(file, "文件信息获取失败", ctx)
	} else {
		response.OkWithDetailed(Res.ExaFileResponse{File: file}, "文件信息获取成功", ctx)
	}
}

func (s *StageApi) DeleteFile(ctx *gin.Context) {
	fileR := Req.GetFileRequest{}
	err := ctx.ShouldBindJSON(&fileR)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := stageService.DeleteFile(fileR.FileId); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}
