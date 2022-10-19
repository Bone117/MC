package model

import "server/global"

type Stage struct {
	global.MODEL
	StageName        string `json:"stage_name" gorm:"comment:比赛阶段"` // 比赛阶段
	CompetitionTimes []CompetitionTime
}
