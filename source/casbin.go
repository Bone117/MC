package source

//
//import (
//	"context"
//	"errors"
//	"server/service"
//
//	adapter "github.com/casbin/gorm-adapter/v3"
//	"gorm.io/gorm"
//)
//
//type initCasbin struct {
//}
//
//func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return ctx, service.ErrMissingDBContext
//	}
//	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
//}
//
//// TableCreated 检查表是否存在
//func (i *initCasbin) TableCreated(ctx context.Context) bool {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return false
//	}
//	return db.Migrator().HasTable(&adapter.CasbinRule{})
//}
//
//func (i initCasbin) InitializerName() string {
//	var entity adapter.CasbinRule
//	return entity.TableName()
//}
//
//func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
//	_, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return ctx, service.ErrMissingDBContext
//	}
//	return ctx, nil
//	//entities := []adapter.CasbinRule{}
//}
//
//func (i *initCasbin) DataInserted(ctx context.Context) bool {
//	db, ok := ctx.Value("db").(*gorm.DB)
//	if !ok {
//		return false
//	}
//	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "777", V1: "GET", V2: "/user/getUserInfo"}).First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) {
//		return false
//	}
//	return true
//}
