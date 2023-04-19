package middleware

import (
	"net/http"
	"server/global"
	"server/model/common/response"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CheckStage() gin.HandlerFunc {
	//TODO 中间件还能改
	return func(context *gin.Context) {
		var stage int
		now := time.Now()
		err := global.DB.Debug().Table("competition_times").Select("stage_id").Where("start_time < ? AND end_time > ?", now, now).Scan(&stage).Error
		if err != nil {
			global.LOG.Error("获取届次信息失败!", zap.Error(err))
			response.FailWithMessage("获取届次信息失败!", context)
			context.Abort()
			return
		}

		if stage == 1 && (context.FullPath() == "/stage/upload" || context.FullPath() == "/stage/deleteFile") {
			context.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "You are not allowed to perform this action during the registration stage",
			})
			context.Abort()
			return
		}

		if stage == 2 && (context.FullPath() == "/stage/updateSign" || context.FullPath() == "/stage/sign" || context.FullPath() == "/stage/deleteSign") {
			context.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "You are not allowed to modify registration information during the upload stage",
			})
			context.Abort()
			return
		}

		if (stage == 3 || stage == 4 || stage == 5) && (context.FullPath() == "/stage/updateSign" || context.FullPath() == "/stage/sign" || context.FullPath() == "/stage/deleteSign" || context.FullPath() == "/stage/uploadFile" || context.FullPath() == "/stage/deleteFile") {
			context.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "You are not allowed to modify registration information or files during the review stage",
			})
			context.Abort()
			return
		}

		context.Next()
	}
}
