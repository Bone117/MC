package api

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/response"
	Req "server/model/request"
	Res "server/model/response"
	"server/utils"

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
	jieCiId, err := StageService.GetJieCi()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	sign := &model.Sign{
		//UserId:         utils.GetUserID(ctx),
		UserId:         1,
		WorkName:       signReq.WorkName,
		WorkFileTypeId: signReq.WorkFileTypeId,
		OtherAuthor:    signReq.OtherAuthor,
		WorkAdviser:    signReq.WorkAdviser,
		WorkDesc:       signReq.WorkDesc,
		JieCiId:        jieCiId,
	}
	stu := &model.Student{
		UserId: 1,
		//UserId:    utils.GetUserID(ctx),
		MajorId:   signReq.MajorId,
		GradeName: signReq.GradeName,
	}
	if err := StageService.Sign(*sign, *stu); err != nil {
		global.LOG.Error("报名失败!", zap.Error(err))
		response.FailWithMessage("报名失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("报名成功", ctx)
	}

}

func (s *StageApi) UpdateSign(ctx *gin.Context) {

}

func (s *StageApi) DeleteSign(ctx *gin.Context) {

}

func (s *StageApi) GetSign(ctx *gin.Context) {

}

func (s *StageApi) GetSignList(ctx *gin.Context) {

}

func (s *StageApi) UploadFile(ctx *gin.Context) {
	//var file model.File
	//form, err := ctx.MultipartForm()
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}

	userid := utils.GetUserID(ctx)
	filePath, key, uploadErr := utils.UploadFile(header, userid) // 文件上传后拿到文件路径
	if uploadErr != nil {
		panic(err)
	}
	fmt.Println(key)
	f1 := model.File{
		Url:      filePath,
		FileName: header.Filename,
	}
	fmt.Println(f1)

	//files := form.File["files"]
	//for _, postFile := range files {
	//dst := path.Join("./", postFile.Filename)
	//err := ctx.SaveUploadedFile(postFile, dst)
	//if err != nil {
	//	global.LOG.Error("文件写入错误!", zap.Error(err))
	//	response.OkWithDetailed(Res.ExaFileResponse{File: file}, "上传失败", ctx)
	//	return
	//}

	response.OkWithDetailed(Res.ExaFileResponse{File: f1}, "上传成功", ctx)
}
