package model

import "server/global"

type Sign struct {
	global.MODEL
	UserId         uint   `json:"userId" gorm:"not null;comment:用户id"`         // 用户id
	JieCiId        uint   `json:"jieCiId"`                                     // 届次
	WorkName       string `json:"workName" gorm:"not null;comment:作品名称"`       // 作品名称
	WorkFileTypeId uint   `json:"workFileTypeId" gorm:"not null;comment:作品类型"` // 作品类型
	OtherAuthor    string `json:"otherAuthor" gorm:"comment:其他作者"`             // 其他作者
	WorkAdviser    string `json:"workAdviser" gorm:"comment:指导老师"`             // 指导老师
	WorkSoftware   string `json:"workSoftware" gorm:"comment:平台"`              // 平台
	IsSubmit       bool   `json:"isSubmit" gorm:"default:false;comment:是否提交"`  //是否提交
	WorkDesc       string `json:"workDesc" gorm:"not null;comment:作品简介"`       // 作品简介
}
