package model

import "server/global"

type Student struct {
	global.MODEL
	UserId    uint   `json:"userId" gorm:"uniqueIndex:user_grade;not null;comment:用户id"`    // 用户id
	MajorId   uint   `json:"majorId" gorm:"not null;comment:专业id"`                          // 专业id
	GradeId   uint   `json:"gradeId" gorm:"not null;comment:班级id"`                          // 班级id
	GradeName string `json:"gradeName" gorm:"uniqueIndex:user_grade;not null;comment:班级名称"` // 班级名称
}
