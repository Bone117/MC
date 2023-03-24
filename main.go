package main

import (
	"server/core"
	"server/global"
	"server/initialize"

	"go.uber.org/zap"
)

func main() {
	//	加载配置
	global.VP = core.Viper() // 初始化Viper
	global.LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.LOG)
	//	初始化mysql
	// CREATE DATABASE m_competition CHARACTER SET utf8mb4
	global.DB = initialize.Gorm() // gorm连接数据库
	if global.DB != nil {
		initialize.RegisterTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
