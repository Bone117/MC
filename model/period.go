package model

import "server/global"

type Period struct {
	global.MODEL
	JieCi            string `json:"jieCi" gorm:"unique;not null;comment:比赛届次"` // 比赛届次
	CompetitionTimes []CompetitionTime
}
