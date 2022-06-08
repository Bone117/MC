package global

import (
	"server/config"
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	DBList map[string]*gorm.DB
	// REDIS  *redis.Client
	CONFIG config.Server
	VP     *viper.Viper
	// LOG    *oplogging.Logger
	LOG *zap.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	Concurrency_Control = &singleflight.Group{}
	lock                sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
