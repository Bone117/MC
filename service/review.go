package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
)

type ReviewService struct {
}

func (r *ReviewService) CreateReview(review model.ReviewSign) error {
	//err := global.DB.Create(&review).Error
	//return err
	// 查询是否已经存在相同的记录
	var existingReview model.ReviewSign
	err := global.DB.Where(&model.ReviewSign{
		ReviewUserId: review.ReviewUserId,
		SignId:       review.SignId,
	}).FirstOrCreate(&existingReview).Error
	if err != nil {
		return err
	}

	// 如果存在相同的记录则更新，否则创建新记录
	if existingReview.ID > 0 {
		err = global.DB.Model(&existingReview).Updates(&review).Error
		if err != nil {
			return err
		}
	} else {
		err = global.DB.Create(&review).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ReviewService) DeleteReview(reviewId uint) error {
	err := global.DB.Where("id=?", reviewId).Delete(&model.ReviewSign{}).Error
	return err
}

func (r *ReviewService) UpdateReview(review model.ReviewSign) error {
	return global.DB.Updates(&review).Error
}

func (r *ReviewService) GetReviewList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var reviewSignList []model.ReviewSign
	var signList []model.Sign
	limit := pageInfo.PageSize
	offset := pageInfo.Page

	db := global.DB.Model(&model.ReviewSign{})
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
		db = db.Where(whereStr, whereArgs...)
	}

	// 获取总数
	if err = db.Count(&total).Error; err != nil {
		return
	}

	// 获取ReviewSign列表
	if err = db.Limit(limit).Offset(offset).Find(&reviewSignList).Error; err != nil {
		return
	}

	// 获取Sign列表
	for _, reviewSign := range reviewSignList {
		sign := model.Sign{}
		err = global.DB.Debug().Where("id = ?", reviewSign.SignId).Preload("Files").First(&sign).Error
		if err != nil {
			return
		}
		signList = append(signList, sign)
	}

	return signList, total, nil
}
