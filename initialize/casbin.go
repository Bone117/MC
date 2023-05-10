package initialize

import (
	"server/global"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func SetupCasbin() {
	var count int64
	global.DB.Model(&gormadapter.CasbinRule{}).Count(&count)

	// 如果记录数为0，表示表为空，需要初始化数据
	if count == 0 {
		entities := []gormadapter.CasbinRule{
			{Ptype: "p", V0: "777", V1: "*", V2: "POST"},
			{Ptype: "p", V0: "777", V1: "*", V2: "GET"},
			{Ptype: "p", V0: "777", V1: "*", V2: "DELETE"},
			{Ptype: "p", V0: "666", V1: "/review/getReviewList", V2: "POST"},
			{Ptype: "p", V0: "666", V1: "/stage/updateSign", V2: "POST"},
			{Ptype: "p", V0: "555", V1: "/review/getEvaluateList", V2: "POST"},
			{Ptype: "p", V0: "555", V1: "/review/updateEvaluate", V2: "POST"},
			{Ptype: "p", V0: "555", V1: "/stage/*", V2: "GET"},
			{Ptype: "p", V0: "444", V1: "/stage/*", V2: "POST"},
			{Ptype: "p", V0: "444", V1: "/stage/*", V2: "GET"},
			{Ptype: "p", V0: "444", V1: "/review/createReport", V2: "POST"},
			{Ptype: "p", V0: "666", V1: "/stage/*", V2: "GET"},
			{Ptype: "p", V0: "444", V1: "/stage/review/*", V2: "deny"},
		}
		if err := global.DB.Create(&entities).Error; err != nil {
			global.LOG.Info("Casbin 表 数据初始化失败!")
		} else {
			global.LOG.Info("Casbin 表 数据初始化成功!")
		}
	}

}
