package model

import "server/global"

type Student struct {
	global.MODEL
	UserId    uint   `json:"userId" gorm:"not null;comment:用户id"`     // 用户id
	MajorId   uint   `json:"MajorId" gorm:"not null;comment:专业id"`    // 专业id
	GradeName string `json:"grade_name" gorm:"not null;comment:班级名称"` // 班级名称
}
