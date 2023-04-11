package model

import "server/global"

type Evaluate struct {
	global.MODEL
	EvaluateUserId uint   `json:"userId" gorm:"not null;uniqueIndex:idx_judge_sign_judge_user_id_sign_id;comment:评委id"` // 评委id
	SignId         uint   `json:"signId" gorm:"not null;uniqueIndex:idx_judge_sign_judge_user_id_sign_id;comment:作品id"` // 作品id
	JieCiId        uint   `json:"jieCiId"`                                                                              // 届次
	Score          uint   `json:"score" gorm:"score:分数"`
	Comments       string `json:"comments" gorm:"comments:评价"`
}
