package model

import (
	"server/global"
	"time"
)

type CompetitionTime struct {
	global.MODEL
	Period    uint      `json:"period" gorm:"index:idx_stage_period,uniqueIndex:idx_stage_period comment:比赛届次"`
	StageID   uint      `json:"stageID" gorm:"index:idx_stage_period,uniqueIndex:idx_stage_period comment:比赛阶段"`
	StartTime time.Time `json:"startTime" gorm:"comment:比赛开始时间"`
	EndTime   time.Time `json:"endTime" gorm:"comment:比赛结束时间"`
}
