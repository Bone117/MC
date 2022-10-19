package model

import "server/global"

type FileType struct {
	global.MODEL
	FileTypeName string `json:"file_type_name" gorm:"comment:文件类型"` // 文件类型
}
