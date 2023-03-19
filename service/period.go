package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
)

type PeriodService struct {
}

func (s *PeriodService) CreatePeriod(period model.CompetitionTime) error {
	newPeriod := model.CompetitionTime{StartTime: period.StartTime, EndTime: period.EndTime}
	t := model.CompetitionTime{Period: period.Period, StageID: period.StageID}
	result := global.DB.Debug().FirstOrCreate(&period, t)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 0 {
		// 记录已创建
	} else {
		// 记录已存在，更新所有字段
		err := global.DB.Debug().Model(&period).Updates(newPeriod).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PeriodService) UpdatePeriod(period model.Period) error {
	return global.DB.Updates(&period).Error
}

func (s *PeriodService) DeletePeriod(periodID int) error {
	err := global.DB.Where("id=?", periodID).Delete(&model.CompetitionTime{}).Error
	return err
}

func (s *PeriodService) GetPeriodList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var periodList []model.CompetitionTime

	//cpName := reflect.TypeOf(model.CompetitionTime{}).Name()
	//cpTable := fmt.Sprintf(":%ss", cpName)

	limit := pageInfo.PageSize
	//offset := pageInfo.PageSize * (pageInfo.Page - 1)
	offset := pageInfo.Page
	db := global.DB.Model(&model.CompetitionTime{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if pageInfo.Keyword != nil {
		keyWord := pageInfo.Keyword
		whereStr := ""
		whereArgs := []interface{}{}
		for key, val := range keyWord {
			whereStr += fmt.Sprintf("%s = ? ", key)
			whereArgs = append(whereArgs, val)
			if len(whereArgs) != len(keyWord) {
				whereStr += "AND "
			}
		}
		err = db.Debug().Limit(limit).Offset(offset).Where(whereStr, whereArgs...).Find(&periodList).Statement.Error
	} else {
		err = db.Debug().Limit(limit).Offset(offset).Find(&periodList).Error
	}

	return periodList, total, err
}

func (s *PeriodService) GetPeriod(periodId int) (model.Period, error) {
	var period model.Period
	err := global.DB.Where("id", periodId).Preload("CompetitionTimes").First(&period).Error
	return period, err
}
