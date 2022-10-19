package model

import (
	"server/global"
	"time"
)

type CompetitionTime struct {
	global.MODEL
	StageID   uint      `json:"stageID" gorm:"comment:比赛届次"`
	PeriodID  uint      `json:"periodID" gorm:"comment:比赛届次"`
	StartTime time.Time `json:"startTime" gorm:"comment:比赛开始时间"`
	EndTime   time.Time `json:"endTime" gorm:"comment:比赛结束时间"`
}
