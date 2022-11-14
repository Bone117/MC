package model

import "server/global"

type File struct {
	global.MODEL
	FileName   string `json:"fileName" gorm:"comment:文件名"`    // 文件名
	Url        string `json:"url" gorm:"comment:文件地址"`        // 文件地址
	FileTypeID uint   `json:"fileTypeID" gorm:"comment:文件类型"` // 文件类型
}
