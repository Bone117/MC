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
		// 权限表初始化
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

		// 用户表初始化
		err = db.AutoMigrate(
			model.User{})
		if err != nil {
			global.LOG.Error("register User table failed", zap.Error(err))
			os.Exit(0)
		}
		userEntities := []model.User{
			{
				UUID:        uuid.NewV4(),
				Username:    "McAdmin",
				Password:    utils.BcryptHash("McAdmin123"),
				NickName:    "超级管理员",
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

	// 通知表初始化
	err := db.AutoMigrate(
		model.Notice{},
	)
	if err != nil {
		global.LOG.Error("register Notice table failed", zap.Error(err))
		os.Exit(0)
	}

	// 参赛时间
	err = db.AutoMigrate(
		model.CompetitionTime{},
	)
	if err != nil {
		global.LOG.Error("register CompetitionTime table failed", zap.Error(err))
		os.Exit(0)
	}

	// 比赛阶段表初始化
	if exist := global.DB.Migrator().HasTable(&model.Stage{}); !exist {
		err = db.AutoMigrate(model.Stage{})
		if err != nil {
			global.LOG.Error("register Stage table failed", zap.Error(err))
			os.Exit(0)
		}
		stageEntities := []model.Stage{
			{StageName: "报名阶段"},
			{StageName: "上传阶段"},
			{StageName: "审核阶段"},
			{StageName: "评审阶段"},
			{StageName: "展示阶段"},
		}
		if err = global.DB.Create(&stageEntities).Error; err != nil {
			global.LOG.Error("initialize Stage table failed", zap.Error(err))
		}
		global.LOG.Info("initialize Stage table success")
	}

	// 专业表
	if exist := global.DB.Migrator().HasTable(&model.Major{}); !exist {
		err = db.AutoMigrate(
			model.Major{},
		)
		if err != nil {
			global.LOG.Error("register Major table failed", zap.Error(err))
			os.Exit(0)
		}
		majorEntities := []model.Major{
			{MajorName: "教育技术学"},
			{MajorName: "小学教育"},
			{MajorName: "现代教育技术"},
			{MajorName: "应用心理学"},
		}
		if err = global.DB.Create(&majorEntities).Error; err != nil {
			global.LOG.Error("initialize Major table failed", zap.Error(err))
		}
		global.LOG.Info("initialize Major table success")
	}

	// 班级表
	err = db.AutoMigrate(
		model.Grade{},
	)
	if err != nil {
		global.LOG.Error("register Grade table failed", zap.Error(err))
		os.Exit(0)
	}

	// 举报表
	err = db.AutoMigrate(
		model.Report{},
	)
	if err != nil {
		global.LOG.Error("register Report table failed", zap.Error(err))
		os.Exit(0)
	}

	// 报名表初始化
	err = db.AutoMigrate(
		model.Sign{},
	)
	if err != nil {
		global.LOG.Error("register Sign table failed", zap.Error(err))
		os.Exit(0)
	}

	// 作品类型表初始化
	if exist := global.DB.Migrator().HasTable(&model.WorkFileType{}); !exist {
		err = db.AutoMigrate(
			model.WorkFileType{},
		)
		if err != nil {
			global.LOG.Error("register WorkFileType table failed", zap.Error(err))
			os.Exit(0)
		}
		workFileTypeEntities := []model.WorkFileType{
			{WorkFileTypeName: "网站类"},
			{WorkFileTypeName: "课件类"},
			{WorkFileTypeName: "DV类"},
			{WorkFileTypeName: "动画类"},
			{WorkFileTypeName: "平面类"},
			{WorkFileTypeName: "微课"},
			{WorkFileTypeName: "移动应用"},
			{WorkFileTypeName: "创客"},
			{WorkFileTypeName: "展板或空间设计"},
		}
		if err = global.DB.Create(&workFileTypeEntities).Error; err != nil {
			global.LOG.Error("initialize WorkFileType table failed", zap.Error(err))
		}
		global.LOG.Info("initialize WorkFileType table success")
	}

	// 文件类型表初始化
	if exist := global.DB.Migrator().HasTable(&model.FileType{}); !exist {
		err = db.AutoMigrate(
			model.FileType{},
		)
		if err != nil {
			global.LOG.Error("register FileType table failed", zap.Error(err))
			os.Exit(0)
		}
		fileTypeEntities := []model.FileType{
			{FileTypeName: "源文件"},
			{FileTypeName: "运行文件"},
			{FileTypeName: "演示文件"},
		}
		if err = global.DB.Create(&fileTypeEntities).Error; err != nil {
			global.LOG.Error("initialize FileType table failed", zap.Error(err))
		}
		global.LOG.Info("initialize FileType table success")
	}

	// 文件表初始化
	err = db.AutoMigrate(
		model.File{},
	)
	if err != nil {
		global.LOG.Error("register File table failed", zap.Error(err))
		os.Exit(0)
	}

	// 作品集锦表初始化
	err = db.AutoMigrate(
		model.Portfolio{},
	)
	if err != nil {
		global.LOG.Error("register Portfolio table failed", zap.Error(err))
		os.Exit(0)
	}

	// 点赞表
	err = db.AutoMigrate(
		model.UserLike{},
	)
	if err != nil {
		global.LOG.Error("register UserLike table failed", zap.Error(err))
		os.Exit(0)
	}

	// 审核表初始化
	err = db.AutoMigrate(
		model.ReviewSign{},
	)
	if err != nil {
		global.LOG.Error("register review table failed", zap.Error(err))
		os.Exit(0)
	}

	// 评委表初始化
	err = db.AutoMigrate(
		model.Evaluate{},
	)
	if err != nil {
		global.LOG.Error("register evaluate table failed", zap.Error(err))
		os.Exit(0)
	}
}
