package initialize

import (
	"os"
	"server/global"
	"server/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Authority{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
