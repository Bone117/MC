package model

import "server/global"

type Notice struct {
	global.MODEL
	Title   string `json:"title" gorm:"type:varchar(250) not null;default:''"`
	Desc    string `json:"desc" gorm:"type:varchar(250) not null;default:''"`
	Content string `json:"content" gorm:"type:longtext default null"`
}
