package service

import (
	"server/global"
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinService struct {
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.CONFIG.Casbin.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
