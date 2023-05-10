package model

import "server/global"

type Portfolio struct {
	global.MODEL
	Title    string `json:"title" gorm:"comment:标题"`     // 标题
	FileName string `json:"fileName" gorm:"comment:文件名"` // 文件名
	Url      string `json:"url" gorm:"comment:文件地址"`     // 文件地址
	CoverUrl string `json:"-"`                           // 缩略图
}
