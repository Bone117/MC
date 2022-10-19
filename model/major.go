package model

import "server/global"

type Major struct {
	global.MODEL
	MajorName string `json:"major_name" gorm:"not null;comment:专业名称"` // 专业名称
	Grades    []Grade
}
