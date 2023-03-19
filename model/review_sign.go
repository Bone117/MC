package model

import "server/global"

type ReviewSign struct {
	global.MODEL
	ReviewUserId uint `json:"userId" gorm:"not null;uniqueIndex:idx_review_sign_review_user_id_sign_id;comment:审核用户id"` // 审核用户id
	SignId       uint `json:"signId" gorm:"not null;uniqueIndex:idx_review_sign_review_user_id_sign_id;comment:作品id"`   // 作品id
	JieCiId      uint `json:"jieCiId"`                                                                                  // 届次
	Sign         Sign `gorm:"foreignkey:SignId;references:ID"`
}
