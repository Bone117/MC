package model

import "server/global"

type Grade struct {
	global.MODEL
	GradeName string `json:"gradeName" gorm:"uniqueIndex:gradename_majorid;not null;comment:班级名称"` // 班级名称
	MajorId   uint   `json:"majorId" gorm:"uniqueIndex:gradename_majorid; comment:专业id"`
	Students  []Student
}
