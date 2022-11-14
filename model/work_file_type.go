package model

import "server/global"

type WorkFileType struct {
	global.MODEL
	WorkFileTypeName string `json:"workFileTypeName" gorm:"comment:作品类型"` // 作品类型
}
