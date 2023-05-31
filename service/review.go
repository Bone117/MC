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

//func (r *ReviewService) UpdateSignStatusByID(id uint,status uint) error {
//	err := global.DB.Model(&model.Sign{}).
//		Where("id = ?", id).
//		Updates(map[string]interface{}{"status": status}).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (r *ReviewService) CreateOrUpdateReviews(reviews []*model.ReviewSign, signIds []uint, userIds []uint) error {
	// 使用事务处理批量创建或更新操作
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除传入的SignId和UserId组合之外的记录
		if len(userIds) == 0 {
			// 当userIds为空时，删除所有userId对应的数据
			err := tx.Unscoped().Where("sign_id IN (?)", signIds).Delete(&model.ReviewSign{}).Error
			if err != nil {
				return err
			}
		} else {
			// 当userIds非空时，按照原逻辑处理
			err := tx.Unscoped().Where("sign_id IN (?)", signIds).Not("review_user_id IN (?)", userIds).Delete(&model.ReviewSign{}).Error
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
		// 修改sign的状态
		// 删除传入的SignId和UserId组合之外的记录
		if len(userIds) == 0 {
			// 当userIds为空时，删除所有userId对应的数据
			err := tx.Unscoped().Where("sign_id IN (?)", signIds).Delete(&model.Evaluate{}).Error
			if err != nil {
				return err
			}
		} else {
			// 当userIds非空时，按照原逻辑处理
			err := tx.Debug().Unscoped().Where("sign_id IN (?)", signIds).Not("evaluate_user_id IN (?)", userIds).Delete(&model.Evaluate{}).Error
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
	//return global.DB.Debug().Model(&model.Evaluate{}).Where("evaluate_user_id = ? AND sign_id = ?", evaluate.EvaluateUserId, evaluate.SignId).Updates(&evaluate).Omit("jie_ci_id").Error
	// 初始化事务
	tx := global.DB.Begin()

	// 更新Sign的status为3
	if err := tx.Model(&model.Sign{}).Where("id = ?", evaluate.SignId).Update("status", 3).Error; err != nil {
		tx.Rollback() // 出错时回滚事务
		return err
	}

	// 更新Evaluate
	if err := tx.Model(&model.Evaluate{}).Where("evaluate_user_id = ? AND sign_id = ?", evaluate.EvaluateUserId, evaluate.SignId).Updates(&evaluate).Omit("jie_ci_id").Error; err != nil {
		tx.Rollback() // 出错时回滚事务
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func (r *ReviewService) GetReviewList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var reviewSignList []model.ReviewSign
	var signList []model.SignWithPhone
	limit := pageInfo.PageSize
	offset := pageInfo.Page

	db := global.DB.Model(&model.ReviewSign{})
	if pageInfo.Keyword != nil {
		keyWord := pageInfo.Keyword
		whereStr := ""
		whereArgs := []interface{}{}
		for key, val := range keyWord {
			fieldName := strcase.ToSnake(key)
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
	if err = db.Limit(limit).Offset(offset).Find(&reviewSignList).Error; err != nil {
		return
	}

	// 获取Sign列表
	for _, reviewSign := range reviewSignList {
		var signWithPhone model.SignWithPhone
		err = global.DB.
			Table("signs").
			Select("signs.*, users.phone").
			Joins("JOIN users ON users.id = signs.user_id").
			Where("signs.id = ?", reviewSign.SignId).
			Preload("Files").
			Scan(&signWithPhone).Error // 使用Scan方法映射到SignWithPhone结构体
		if err != nil {
			return
		}
		//sign := signWithPhone.Sign
		//sign.Phone = signWithPhone.Phone // 将Phone字段值赋给Sign模型
		signList = append(signList, signWithPhone)
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
		err = global.DB.Where("id = ?", reviewSign.SignId).Preload("Files").Preload("Evaluates", "evaluate_user_id = ?", userID).First(&sign).Error
		if err != nil {
			return
		}
		signList = append(signList, sign)
	}
	return signList, total, nil
}

func (r *ReviewService) GetSignWithEvaluateList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var SignList []model.Sign
	limit := pageInfo.PageSize
	offset := pageInfo.Page

	db := global.DB.Model(&model.Sign{})
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

	// 获取Sign列表
	if err = db.Limit(limit).Offset(offset).Preload("Evaluates.User").Find(&SignList).Error; err != nil {
		return
	}

	return SignList, total, nil
}

func (r *ReviewService) CreateOrUpdateReport(report model.Report) error {
	var existingReport model.Report

	// 查找是否已存在具有相同 ReportUserId 和 SignId 的记录
	result := global.DB.Where("report_user_id = ? AND sign_id = ?", report.ReportUserId, report.SignId).First(&existingReport)

	if result.Error == nil {
		// 如果找到记录，则更新 Comments 字段
		existingReport.Content = report.Content
		return global.DB.Save(&existingReport).Error
	} else if result.Error == gorm.ErrRecordNotFound {
		// 如果没有找到记录，则创建新的记录
		return global.DB.Create(&report).Error
	} else {
		// 其他错误
		return result.Error
	}
}

//func (r *ReviewService) GetReportList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
//	var reportList []model.Report
//	limit := pageInfo.PageSize
//	offset := pageInfo.Page
//
//	db := global.DB.Model(&model.Report{})
//	if pageInfo.Keyword != nil {
//		keyWord := pageInfo.Keyword
//		whereStr := ""
//		whereArgs := []interface{}{}
//		for key, val := range keyWord {
//			whereStr += fmt.Sprintf("%s = ? ", key)
//			whereArgs = append(whereArgs, val)
//			if len(whereArgs) != len(keyWord) {
//				whereStr += "AND "
//			}
//		}
//		db = db.Where(whereStr, whereArgs...)
//	}
//
//	// 获取总数
//	if err = db.Count(&total).Error; err != nil {
//		return
//	}
//
//	// 获取Report列表
//	if err = db.Limit(limit).Offset(offset).Find(&reportList).Error; err != nil {
//		return
//	}
//
//	signList := []model.Sign{}
//
//	// 获取Sign列表
//	for _, report := range reportList {
//		sign := model.Sign{}
//		err = global.DB.Debug().Where("id = ?", report.SignId).First(&sign).Error
//		if err != nil {
//			return
//		}
//		sign.Reports = report
//		signList = append(signList, sign)
//	}
//
//	return signList, total, nil
//}

func (r *ReviewService) GetReportList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var reportList []model.Report
	limit := pageInfo.PageSize
	offset := pageInfo.Page
	db := global.DB.Model(&model.Report{})
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
	if err = db.Count(&total).Error; err != nil {
		return
	}
	// 获取Report列表
	if err = db.Limit(limit).Offset(offset).Find(&reportList).Error; err != nil {
		return
	}
	signList := []model.Sign{}
	// 获取Sign列表
	for _, report := range reportList {
		sign := model.Sign{}
		err = global.DB.Where("id = ?", report.SignId).First(&sign).Error
		if err != nil {
			return
		}

		// 获取用户列表
		var reportWithUsername model.ReportWithUsername
		err = global.DB.Table("reports").
			Select("reports.*, users.username").
			Joins("LEFT JOIN users ON users.id = reports.report_user_id").
			Where("reports.id = ?", report.ID).
			First(&reportWithUsername).Error

		if err != nil {
			return
		}

		// 将带有用户名的举报添加到签名
		sign.ReportWithUsername = reportWithUsername

		signList = append(signList, sign)
	}
	return signList, total, nil
}
