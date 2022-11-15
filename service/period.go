package service

import (
	"server/global"
	"server/model"
	"server/model/common/request"
)

type PeriodService struct {
}

func (s *PeriodService) CreatePeriod(period model.Period) error {
	return global.DB.Create(&period).Error
}

func (s *PeriodService) UpdatePeriod(period model.Period) error {
	return global.DB.Updates(&period).Error
}

func (s *PeriodService) DeletePeriod(periodID int) error {
	err := global.DB.Where("id=?", periodID).Delete(&model.Period{}).Error
	return err
}

func (s *PeriodService) GetPeriodList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var periodList []model.Period

	//cpName := reflect.TypeOf(model.CompetitionTime{}).Name()
	//cpTable := fmt.Sprintf(":%ss", cpName)
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.DB.Model(&model.Period{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Preload("CompetitionTimes").Find(&periodList).Error

	return periodList, total, err
}

func (s *PeriodService) GetPeriod(periodId int) (model.Period, error) {
	var period model.Period
	err := global.DB.Where("id", periodId).Preload("CompetitionTimes").First(&period).Error
	return period, err
}
