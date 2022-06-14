package api

import (
	"server/global"
	"server/model"
	"server/model/common/response"
	Res "server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityApi struct {
}

func (a *AuthorityApi) CreateAuthority(ctx *gin.Context) {
	var authority model.Authority
	_ = ctx.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if authBack, err := authorityService.CreateAuthority(authority); err != nil {
		global.LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), ctx)
	} else {
		//_ = menuService.AddMenuAuthority(Req.DefaultMenu(), authority.AuthorityId)
		response.OkWithDetailed(Res.AuthorityResponse{Authority: authBack}, "创建成功", ctx)
	}
}
