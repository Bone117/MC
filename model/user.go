package model

import (
	"server/global"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	global.MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                    // 用户UUID
	Username    string    `json:"username" gorm:"comment:用户登录名"`                 // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                      // 用户登录密码
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
	Authority   Authority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
}
