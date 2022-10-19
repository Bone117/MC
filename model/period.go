package model

import "server/global"

type Period struct {
	global.MODEL
	JieCi            string `json:"jie_ci" gorm:"comment:比赛届次"` // 比赛届次
	CompetitionTimes []CompetitionTime
}
