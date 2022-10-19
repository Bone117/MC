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
				{Ptype: "p", V0: "777", V1: "/base/login", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/register", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/user/getUserInfo", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "/user/setUserInfo", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/user/setSelfInfo", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/user/getUserList", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/deleteUser", V2: "DELETE"},
				{Ptype: "p", V0: "777", V1: "/user/changePassword", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/setUserAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/resetPassword", V2: "POST"},

				//{Ptype: "p", V0: "777", V1: "/authority/copyAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/updateAuthority", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/authority/createAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/deleteAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/getAuthorityList", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/menu/getMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getMenuList", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/addBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getBaseMenuTree", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/addMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/deleteBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/updateBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getBaseMenuById", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/findFile", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/breakpointContinueFinish", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/breakpointContinue", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/removeChunk", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/upload", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/casbin/updateCasbin", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
				{Ptype: "p", V0: "777", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

				{Ptype: "p", V0: "777", V1: "/email/emailTest", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/simpleUploader/upload", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/simpleUploader/checkFileMd5", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "/simpleUploader/mergeFileMd5", V2: "GET"},

				{Ptype: "p", V0: "777", V1: "/authorityBtn/setAuthorityBtn", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authorityBtn/getAuthorityBtn", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authorityBtn/canRemoveAuthorityBtn", V2: "POST"},

				{Ptype: "p", V0: "666", V1: "/user/getUserInfo", V2: "GET"},
				{Ptype: "p", V0: "666", V1: "/user/setUserInfo", V2: "PUT"},
				{Ptype: "p", V0: "666", V1: "/user/setSelfInfo", V2: "PUT"},
				{Ptype: "p", V0: "666", V1: "/user/getUserList", V2: "POST"},
				{Ptype: "p", V0: "666", V1: "/user/deleteUser", V2: "DELETE"},
				{Ptype: "p", V0: "666", V1: "/user/changePassword", V2: "POST"},
				{Ptype: "p", V0: "666", V1: "/user/setUserAuthority", V2: "POST"},
				{Ptype: "p", V0: "666", V1: "/user/setUserAuthorities", V2: "POST"},
				{Ptype: "p", V0: "666", V1: "/user/resetPassword", V2: "POST"},

				{Ptype: "p", V0: "9528", V1: "/base/login", V2: "POST"},
				{Ptype: "p", V0: "666", V1: "/user/register", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},

				{Ptype: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},

				{Ptype: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "PUT"},
				{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "GET"},
				{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "DELETE"},
				{Ptype: "p", V0: "9528", V1: "/customer/customerList", V2: "GET"},
				{Ptype: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
				{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"},
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
