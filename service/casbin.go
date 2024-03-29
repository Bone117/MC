package service

import (
	"errors"
	"server/global"
	"server/model/request"
	"sync"

	"github.com/casbin/casbin/v2/model"
	"go.uber.org/zap"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type CasbinService struct {
}

// var (
//
//	syncedEnforcer *casbin.SyncedEnforcer
//	once           sync.Once
//
// )
var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.DB)
		rbacModel := `
			[request_definition]
			r = sub, obj, act
			
			[policy_definition]
			p = sub, obj, act
			
			[role_definition]
			g = _, _
			
			[policy_effect]
			e = some(where (p.eft == allow))
			
			[matchers]
			m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
			`
		m, err := model.NewModelFromString(rbacModel)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}

//func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
//	once.Do(func() {
//		a, _ := gormadapter.NewAdapterByDB(global.DB)
//		syncedEnforcer, _ = casbin.NewSyncedEnforcer(a)
//	})
//	_ = syncedEnforcer.LoadPolicy()
//	return syncedEnforcer
//}

func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// @function: UpdateCasbinApi
// @description: API更新随动
// @param: oldPath string, newPath string, oldMethod string, newMethod string
// @return: error
func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// @function: GetPolicyPathByAuthorityId
// @description: 获取权限列表
// @param: authorityId string
// @return: pathMaps []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}
