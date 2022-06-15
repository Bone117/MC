package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority")
	authorityApi := api.ApiGroupApp.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority) // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
		authorityRouter.PUT("updateAuthority", authorityApi.UpdateAuthority)  // 更新角色
		//authorityRouter.POST("setDataAuthority", authorityApi.SetDataAuthority) // 设置角色资源权限
		authorityRouter.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}

}
