package model

import "server/global"

type Grade struct {
	global.MODEL
	GradeName string `json:"gradeName" gorm:"not null;comment:班级名称"` // 班级名称
	MajorId   uint   `json:"majorId" gorm:"not null comment:专业id"`
	UserId    uint   `json:"userId" gorm:"uniqueIndex not null;comment:用户id"` // 用户id
}
