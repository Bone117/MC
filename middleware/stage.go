package middleware

import (
	"net/http"
	"server/global"
	"server/utils"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CheckStage() gin.HandlerFunc {
	//TODO 中间件还能改
	return func(context *gin.Context) {
		userid := utils.GetUserID(context)
		var authority int
		err := global.DB.Debug().Table("user_authority").Select("authority_authority_id").Where("user_id = ?", userid).Scan(&authority).Error
		if err != nil {
			global.LOG.Error("获取权限出错!", zap.Error(err))
			context.Abort()
			return
		}
		if authority == 777 {
			context.Next()
			return
		}

		var stage int
		now := time.Now()
		err = global.DB.Debug().Table("competition_times").Select("stage_id").Where("start_time < ? AND end_time > ? AND deleted_at IS NULL", now, now).Scan(&stage).Error
		if err != nil {
			global.LOG.Error("获取届次信息失败!", zap.Error(err))
			context.Abort()
			return
		}

		if stage >= 3 && (context.FullPath() == "/stage/upload" || context.FullPath() == "/stage/deleteFile") {
			context.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "You are not allowed to perform this action during the registration stage",
			})
			context.Abort()
			return
		}

		if stage >= 3 && (context.FullPath() == "/stage/sign" || context.FullPath() == "/stage/deleteSign") {
			context.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "You are not allowed to modify registration information during the upload stage",
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
