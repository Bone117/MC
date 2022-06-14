package initialize

import (
	"server/global"
	"server/model"
	"server/utils"
	"sync"

	uuid "github.com/satori/go.uuid"

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

				{Ptype: "p", V0: "777", V1: "/authority/copyAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/updateAuthority", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/authority/createAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/deleteAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/getAuthorityList", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/authority/setDataAuthority", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/menu/getMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getMenuList", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/addBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getBaseMenuTree", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/addMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/deleteBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/updateBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/menu/getBaseMenuById", V2: "POST"},

				{Ptype: "p", V0: "777", V1: "/user/getUserInfo", V2: "GET"},
				{Ptype: "p", V0: "777", V1: "/user/setUserInfo", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/user/setSelfInfo", V2: "PUT"},
				{Ptype: "p", V0: "777", V1: "/user/getUserList", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/deleteUser", V2: "DELETE"},
				{Ptype: "p", V0: "777", V1: "/user/changePassword", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/setUserAuthority", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/setUserAuthorities", V2: "POST"},
				{Ptype: "p", V0: "777", V1: "/user/resetPassword", V2: "POST"},

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

				{Ptype: "p", V0: "7771", V1: "/base/login", V2: "POST"},
				{Ptype: "p", V0: "555", V1: "/user/register", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/api/createApi", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/api/getApiList", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/api/getApiById", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/api/deleteApi", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/api/updateApi", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/api/getAllApis", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/authority/createAuthority", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/authority/deleteAuthority", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/authority/getAuthorityList", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/authority/setDataAuthority", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/getMenu", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/getMenuList", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/addBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/getBaseMenuTree", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/addMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/getMenuAuthority", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/deleteBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/updateBaseMenu", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/menu/getBaseMenuById", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/user/changePassword", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/user/getUserList", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/user/setUserAuthority", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/fileUploadAndDownload/upload", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/fileUploadAndDownload/editFileName", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/casbin/updateCasbin", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/jwt/jsonInBlacklist", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/system/getSystemConfig", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/system/setSystemConfig", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/customer/customer", V2: "POST"},
				{Ptype: "p", V0: "7771", V1: "/customer/customer", V2: "PUT"},
				{Ptype: "p", V0: "7771", V1: "/customer/customer", V2: "DELETE"},
				{Ptype: "p", V0: "7771", V1: "/customer/customer", V2: "GET"},
				{Ptype: "p", V0: "7771", V1: "/customer/customerList", V2: "GET"},
				{Ptype: "p", V0: "7771", V1: "/user/getUserInfo", V2: "GET"},

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
					global.LOG.Error("Casbin 表 数据初始化失败!")
				} else {
					global.LOG.Info("Casbin 表 数据初始化成功!")
				}
			}

		}
		InitializeAuthority()
	})

}

func InitializeAuthority() {
	authorityEntities := []model.Authority{
		{AuthorityId: "777", AuthorityName: "超级管理员"},
		{AuthorityId: "666", AuthorityName: "审核员"},
		{AuthorityId: "555", AuthorityName: "评委"},
		{AuthorityId: "444", AuthorityName: "学生"},
	}
	if err := global.DB.Create(&authorityEntities).Error; err != nil {
		global.LOG.Error("authority表初始化失败")
	}
	userEntities := []model.User{
		{
			UUID:        uuid.NewV4(),
			Username:    "admin",
			Password:    utils.BcryptHash("admin123"),
			NickName:    "超级管理员1",
			Authorities: []model.Authority{{AuthorityId: "777"}},
			Phone:       "17857094799",
			Email:       "785484564@qq.com",
		},
		{
			UUID:        uuid.NewV4(),
			Username:    "shy",
			Password:    utils.BcryptHash("test123"),
			NickName:    "审核员1",
			Authorities: []model.Authority{{AuthorityId: "666"}},
			Phone:       "17611111111",
			Email:       "333333333@qq.com"},
		{
			UUID:        uuid.NewV4(),
			Username:    "pw",
			Password:    utils.BcryptHash("test123"),
			NickName:    "评委1",
			Authorities: []model.Authority{{AuthorityId: "555"}},
			Phone:       "17611111111",
			Email:       "333333333@qq.com"},
		{
			UUID:        uuid.NewV4(),
			Username:    "student",
			Password:    utils.BcryptHash("test123"),
			NickName:    "学生1",
			Authorities: []model.Authority{{AuthorityId: "444"}},
			Phone:       "17611111111",
			Email:       "333333333@qq.com"},
	}
	if err := global.DB.Create(&userEntities).Error; err != nil {
		global.LOG.Error("user表初始化失败")
	}

	//userAuthority := []model.UseAuthority{
	//	{UserId: 1, AuthorityAuthorityId: "777"},
	//	{UserId: 2, AuthorityAuthorityId: "666"},
	//	{UserId: 3, AuthorityAuthorityId: "555"},
	//	{UserId: 4, AuthorityAuthorityId: "444"},
	//}
	//if err := global.DB.Create(&userAuthority).Error; err != nil {
	//	global.LOG.Error("authority表初始化失败")
	//}
}
