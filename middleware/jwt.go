package middleware

import (
	"server/global"
	"server/model/common/response"
	"server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", context)
			context.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", context)
				context.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), context)
			context.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			context.Header("new-token", newToken)
			context.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		context.Set("claims", claims)
		context.Next()
	}
}
