package service

import (
	"errors"
	"fmt"
	"server/global"
	"server/model"
	"strconv"
	"time"
)

type StageService struct{}

func (s *StageService) Sign(sign model.Sign, mg model.Student) error {
	// 班级表
	grade := model.Grade{GradeName: mg.GradeName, MajorId: mg.MajorId}
	rowAffected := global.DB.Where("grade_name = ? And major_id = ?", mg.GradeName, mg.MajorId).First(&grade).RowsAffected
	if rowAffected == 0 {
		err := global.DB.Create(&grade).Error
		if err != nil {
			return errors.New("grade create failed")
		}
	}
	// 学生表
	rowAffected = global.DB.Where("grade_name = ? And user_id = ?", mg.GradeName, mg.UserId).First(&mg).RowsAffected
	if rowAffected == 0 {
		err := global.DB.Create(&mg).Error
		if err != nil {
			return errors.New("student create failed")
		}
	}
	// 报名表
	var count int64
	global.DB.Model(&model.Sign{}).Where("user_id = ?", sign.UserId).Count(&count)
	if count > global.CONFIG.Sign.Number {
		errMsg := "不能超出" + strconv.FormatInt(global.CONFIG.Sign.Number, 10) + "项报名内容"
		return errors.New(errMsg)
	}
	return global.DB.Create(&sign).Error
}

func (s StageService) GetGrade(gradeName string) (uint, error) {
	grade := model.Grade{}
	err := global.DB.Where("grade_name = ?", gradeName).First(&grade).Error
	return grade.ID, err
}

func (s *StageService) GetJieCi() (uint, error) {
	cpTime := model.CompetitionTime{}
	now := time.Now()
	fmt.Println("The time is", now)
	err := global.DB.Where("start_time < ? AND end_time > ?", now, now).First(&cpTime).Error
	return cpTime.ID, err
}

func (s *StageService) Upload(file model.File) error {
	return global.DB.Create(&file).Error
}
