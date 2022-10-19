package model

import "server/global"

type WorkFileType struct {
	global.MODEL
	WorkFileTypeName string `json:"work_file_type_name" gorm:"comment:作品类型"` // 作品类型
}
