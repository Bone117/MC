package model

import "server/global"

type Grade struct {
	global.MODEL
	GradeName string `json:"grade_name" gorm:"not null;comment:班级名称"` // 班级名称
}
