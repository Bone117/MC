package model

import (
	"server/global"
)

type UserLike struct {
	global.MODEL
	UserID uint `json:"userId" gorm:"not null;comment:用户id"`
	SignID uint `json:"signId" gorm:"not null;comment:作品id"`
}
