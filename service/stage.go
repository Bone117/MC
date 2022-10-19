package service

import (
	"fmt"
	"server/global"
	"server/model"
	"time"
)

type StageService struct{}

func (s *StageService) Sign(sign model.Sign) error {
	return global.DB.Create(&sign).Error
}

func (s *StageService) GetJieCi() (uint, error) {
	cpTime := model.CompetitionTime{}
	now := time.Now()
	fmt.Println("The time is", now)
	err := global.DB.Where("start_time < ? AND end_time > ?", now, now).Find(&cpTime).Error
	return cpTime.ID, err
}

func (s *StageService) Upload(file model.File) error {
	return global.DB.Create(&file).Error
}
