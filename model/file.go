package model

import "server/global"

type File struct {
	global.MODEL
	FileName   string `json:"FileName" gorm:"comment:文件名"`      // 文件名
	Url        string `json:"url" gorm:"comment:文件地址"`          // 文件地址
	FileTypeID uint   `json:"file_type_id" gorm:"comment:文件类型"` // 文件类型
}