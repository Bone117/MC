package model

import "server/global"

type Report struct {
	global.MODEL
	ReportUserId uint   `json:"userId" gorm:"not null;uniqueIndex:idx_judge_sign_judge_user_id_sign_id;comment:举报用户id"` // 举报用户id
	SignId       uint   `json:"signId" gorm:"not null;uniqueIndex:idx_judge_sign_judge_user_id_sign_id;comment:作品id"`   // 作品id
	JieCiId      uint   `json:"jieCiId"`                                                                                // 届次
	Content      string `json:"content"`
}
