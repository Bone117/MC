package core

import (
	"fmt"
	"server/global"
	"server/initialize"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunServer() {
	//if global.CONFIG.System.UseMultipoint || global.CONFIG.System.UseRedis {
	//	// 初始化redis服务
	//	initialize.Redis()
	//}

	// 从db加载jwt数据
	//if global.DB != nil {
	//system.LoadAll()
	//}

	// 设置casbin
	initialize.SetupCasbin()

	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	global.LOG.Error(s.ListenAndServe().Error())
}
