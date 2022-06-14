package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	baseApi := api.ApiGroupApp.BaseApi
	{
		userRouter.POST("register", baseApi.Register)             // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword) // 用户修改密码
		userRouter.POST("resetPassword", baseApi.ResetPassword)   // 重置密码
		//userRouter.POST("setUserAuthority", baseApi.SetUserAuthority) // 设置用户权限
		userRouter.DELETE("deleteUser", baseApi.DeleteUser) // 删除用户
		userRouter.PUT("setUserInfo", baseApi.SetUserInfo)  // 设置用户信息
		userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)  // 设置自身信息
		//userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限
	}
}
