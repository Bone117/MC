package api

import (
	"server/global"
	"server/model/common/response"
	"server/model/request"
	Res "server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CasbinApi struct{}

func (cas *CasbinApi) UpdateCasbin(ctx *gin.Context) {
	var cmr request.CasbinInReceive
	_ = ctx.ShouldBindJSON(&cmr)
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{data=systemRes.PolicyPathResponse,msg=string} "获取权限列表,返回包括casbin详情列表"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(ctx *gin.Context) {
	var casbin request.CasbinInReceive
	_ = ctx.ShouldBindJSON(&casbin)
	if err := utils.Verify(casbin, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(Res.PolicyPathResponse{Paths: paths}, "获取成功", ctx)
}
