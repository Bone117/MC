package api

import (
	"server/global"
	"server/model/common/response"

	"github.com/gin-gonic/gin"
)

type DBApi struct {
}

func (i *DBApi) InitDB(ctx *gin.Context) {
	if global.DB != nil {
		global.LOG.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", ctx)
		return
	}
	//var dbInfo request.InitDB
}
