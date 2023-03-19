package response

import (
	"server/model"
	"time"

	"gorm.io/gorm"
)

type CpResponse struct {
	StageID   uint      `json:"stageID" gorm:"comment:比赛阶段"`
	StartTime time.Time `json:"startTime" gorm:"comment:比赛开始时间"`
	EndTime   time.Time `json:"endTime" gorm:"comment:比赛结束时间"`
}

type PeriodResponse struct {
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"` // 删除时间
	Period    model.Period
	CpTime    []CpResponse
}
