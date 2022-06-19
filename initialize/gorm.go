package initialize

import (
	"os"
	"server/global"
	"server/model"
	"server/utils"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func RegisterTables(db *gorm.DB) {
	if exist := global.DB.Migrator().HasTable(&model.Authority{}); !exist {
		err := db.AutoMigrate(
			model.Authority{})
		if err != nil {
			global.LOG.Error("register Authority table failed", zap.Error(err))
			os.Exit(0)
		}
		authorityEntities := []model.Authority{
			{AuthorityId: "777", AuthorityName: "超级管理员"},
			{AuthorityId: "666", AuthorityName: "审核员"},
			{AuthorityId: "555", AuthorityName: "评委"},
			{AuthorityId: "444", AuthorityName: "学生"},
		}
		if err := global.DB.Create(&authorityEntities).Error; err != nil {
			global.LOG.Error("register Authority table failed", zap.Error(err))
		}
		global.LOG.Info("register Authority success")

		err = db.AutoMigrate(
			model.User{})
		if err != nil {
			global.LOG.Error("register User table failed", zap.Error(err))
			os.Exit(0)
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
			global.LOG.Error("register User table failed", zap.Error(err))
		}
		global.LOG.Info("register User table success")
	}
	err := db.AutoMigrate(
		model.Notice{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
}
