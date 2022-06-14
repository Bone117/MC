package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.MODEL
	UUID        uuid.UUID   `json:"uuid" gorm:"comment:用户UUID"`                    // 用户UUID
	Username    string      `json:"username" gorm:"not null;unique;comment:用户登录名"` // 用户登录名
	Password    string      `json:"-"  gorm:"comment:用户登录密码"`                      // 用户登录密码
	NickName    string      `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`     // 用户昵称 	// 活跃颜色
	Authorities []Authority `json:"authorities" gorm:"many2many:user_authority"`
	Phone       string      `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email       string      `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}
