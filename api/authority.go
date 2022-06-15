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

type AuthorityApi struct {
}

//TODO CasBin未配置: 每个角色对应权限

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
		_ = casbinService.UpdateCasbin(authority.AuthorityId, Req.DefaultCasbin())
		response.OkWithDetailed(Res.AuthorityResponse{Authority: authBack}, "创建成功", ctx)
	}
}

func (a *AuthorityApi) UpdateAuthority(ctx *gin.Context) {
	var auth model.Authority
	_ = ctx.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	if authority, err := authorityService.UpdateAuthority(auth); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(Res.AuthorityResponse{Authority: authority}, "更新成功", ctx)
	}
}

func (a *AuthorityApi) DeleteAuthority(ctx *gin.Context) {
	var auth model.Authority
	_ = ctx.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 删除角色之前需要判断是否有用户正在使用此角色
	if err := authorityService.DeleteAuthority(&auth); err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}

}

func (a *AuthorityApi) GetAuthorityList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	_ = ctx.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if list, total, err := authorityService.GetAuthorityList(pageInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
