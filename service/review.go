package service

import (
	"errors"
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"

	"github.com/iancoleman/strcase"

	"gorm.io/gorm"
)

type ReviewService struct {
}

func (r *ReviewService) CreateReviews(reviews []*model.ReviewSign, signIds []uint) error {
	// 使用事务处理批量创建或更新操作
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// // 禁用软删除，彻底删除所有指定SignId的现有评估记录
		err := tx.Unscoped().Delete(&model.ReviewSign{}, "sign_id IN (?)", signIds).Error
		if err != nil {
			return err
		}

		// 批量插入新的评估记录
		for _, review := range reviews {
			err = tx.Create(review).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *ReviewService) CreateOrUpdateReviews(reviews []*model.ReviewSign, signIds []uint, userIds []uint) error {
	// 使用事务处理批量创建或更新操作
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除传入的SignId和UserId组合之外的记录
		if len(userIds) == 0 {
			// 当userIds为空时，删除所有userId对应的数据
			err := tx.Unscoped().Debug().Where("sign_id IN (?)", signIds).Delete(&model.ReviewSign{}).Error
			if err != nil {
				return err
			}
		} else {
			// 当userIds非空时，按照原逻辑处理
			err := tx.Unscoped().Debug().Where("sign_id IN (?)", signIds).Not("review_user_id IN (?)", userIds).Delete(&model.ReviewSign{}).Error
			if err != nil {
				return err
			}
		}

		for _, review := range reviews {
			var existingReview model.ReviewSign
			err := tx.Where(&model.ReviewSign{
				ReviewUserId: review.ReviewUserId,
				SignId:       review.SignId,
			}).First(&existingReview).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 如果记录不存在，则创建新记录
				err = tx.Create(review).Error
				if err != nil {
					return err
				}
			} else if err == nil {
				// 如果记录存在，检查是否需要更新
				if review.JieCiId != existingReview.JieCiId {
					err = tx.Model(&existingReview).Updates(review).Error
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		}
		return nil
	})
}

func (r *ReviewService) CreateOrUpdateEvaluates(evaluates []*model.Evaluate, signIds []uint, userIds []uint) error {
	// 使用事务处理批量创建或更新操作
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除传入的SignId和UserId组合之外的记录
		if len(userIds) == 0 {
			// 当userIds为空时，删除所有userId对应的数据
			err := tx.Unscoped().Debug().Where("sign_id IN (?)", signIds).Delete(&model.Evaluate{}).Error
			if err != nil {
				return err
			}
		} else {
			// 当userIds非空时，按照原逻辑处理
			err := tx.Unscoped().Debug().Where("sign_id IN (?)", signIds).Not("evaluate_user_id IN (?)", userIds).Delete(&model.Evaluate{}).Error
			if err != nil {
				return err
			}
		}

		for _, review := range evaluates {
			var existingEvaluate model.Evaluate
			err := tx.Where(&model.Evaluate{
				EvaluateUserId: review.EvaluateUserId,
				SignId:         review.SignId,
			}).First(&existingEvaluate).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 如果记录不存在，则创建新记录
				err = tx.Create(review).Error
				if err != nil {
					return err
				}
			} else if err == nil {
				// 如果记录存在，检查是否需要更新
				if review.JieCiId != existingEvaluate.JieCiId {
					err = tx.Model(&existingEvaluate).Updates(review).Error
					if err != nil {
						return err
					}
				}
			} else {
				return err
			}
		}
		return nil
	})
}

func (r *ReviewService) CreateEvaluates(evaluates []*model.Evaluate) error {
	// 使用事务处理批量创建或更新操作
	return global.DB.Transaction(func(tx *gorm.DB) error {
		for _, evaluate := range evaluates {
			// 查找具有相同EvaluateUserId和SignId的评估
			var existingEvaluate model.Evaluate
			err := tx.Where(&model.Evaluate{
				EvaluateUserId: evaluate.EvaluateUserId,
				SignId:         evaluate.SignId,
			}).First(&existingEvaluate).Error

			// 判断是否找到现有记录
			if err == nil {
				// 如果找到现有记录，删除现有记录
				err = tx.Delete(&existingEvaluate).Error
				if err != nil {
					return err
				}
			}

			// 创建新记录（包括更新现有评委或添加新评委）
			err = tx.Create(evaluate).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *ReviewService) DeleteReview(reviewId uint) error {
	err := global.DB.Where("id=?", reviewId).Delete(&model.ReviewSign{}).Error
	return err
}

func (r *ReviewService) UpdateReview(review model.ReviewSign) error {
	return global.DB.Updates(&review).Error
}
func (r *ReviewService) UpdateEvaluate(evaluate model.Evaluate) error {
	//return global.DB.Debug().Model(&model.Evaluate{}).Where("evaluate_user_id = ? AND sign_id = ?", evaluate.EvaluateUserId, evaluate.SignId).Updates(&evaluate).Select("score", "comments").Error
	return global.DB.Debug().Model(&model.Evaluate{}).Where("evaluate_user_id = ? AND sign_id = ?", evaluate.EvaluateUserId, evaluate.SignId).Updates(&evaluate).Omit("jie_ci_id").Error
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

func (r *ReviewService) GetEvaluateList(pageInfo request.PageInfo, userID uint) (list interface{}, total int64, err error) {
	var EvaluateList []model.Evaluate
	var signList []model.Sign
	limit := pageInfo.PageSize
	offset := pageInfo.Page

	db := global.DB.Model(&model.Evaluate{})
	if pageInfo.Keyword != nil {
		keyWord := pageInfo.Keyword
		whereStr := ""
		whereArgs := []interface{}{}
		for key, val := range keyWord {
			fieldName := strcase.ToSnake(key) // 将小驼峰命名转换为下划线命名
			// 将数字类型的值转换为uint类型
			if v, ok := val.(float64); ok {
				val = uint(v)
			}
			whereStr += fmt.Sprintf("%s = ? ", fieldName)
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
	if err = db.Limit(limit).Offset(offset).Find(&EvaluateList).Error; err != nil {
		return
	}

	// 获取Sign列表
	for _, reviewSign := range EvaluateList {
		sign := model.Sign{}
		err = global.DB.Debug().Where("id = ?", reviewSign.SignId).Preload("Files").Preload("Evaluates", "evaluate_user_id = ?", userID).First(&sign).Error
		if err != nil {
			return
		}
		signList = append(signList, sign)
	}
	return signList, total, nil
}
