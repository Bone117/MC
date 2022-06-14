package middleware

import (
	"server/model/common/response"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

var casbinService = service.ServiceGroupApp.CasbinService

func CasbinHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		waitUse, _ := utils.GetClaims(context)
		// 获取请求的PATH
		obj := context.Request.URL.Path
		// 获取请求方法
		act := context.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			context.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", context)
			context.Abort()
			return
		}
	}
}
