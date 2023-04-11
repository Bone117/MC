package initialize

import (
	"server/global"
	"sync"

	"go.uber.org/zap"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func SetupCasbin() {
	var err error
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		syncedEnforcer, err = casbin.NewSyncedEnforcer(global.CONFIG.Casbin.ModelPath, a)
		if err != nil {
			global.LOG.Error("SetupCasbin failed", zap.Error(err))
		} else {
			_ = syncedEnforcer.LoadPolicy()
			global.LOG.Info("SetupCasbin succeed")
			entities := []gormadapter.CasbinRule{
				{Ptype: "p", V0: "777", V1: "*", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "*", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "*", V2: "DELETE"},
				{Ptype: "p", V0: "666", V1: "/review/getReviewList", V2: "POST"},
				{Ptype: "p", V0: "666", V1: "/stage/updateSign", V2: "POST"},
				{Ptype: "p", V0: "555", V1: "/review/getEvaluateList", V2: "POST"},
				{Ptype: "p", V0: "555", V1: "/review/updateEvaluate", V2: "POST"},
				{Ptype: "p", V0: "444", V1: "/stage/*", V2: "POST"},
				{Ptype: "p", V0: "444", V1: "/stage/*", V2: "GET"},
				{Ptype: "p", V0: "444", V1: "/stage/review/*", V2: "deny"},
			}
			if global.DB.Migrator().HasTable(&gormadapter.CasbinRule{}) {
				if err := global.DB.Create(&entities).Error; err != nil {
					global.LOG.Info("Casbin 表 数据初始化失败!")
				} else {
					global.LOG.Info("Casbin 表 数据初始化成功!")
				}
			}

		}
	})

}
