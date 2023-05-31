package service

import (
	"errors"
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
	"server/utils"
	"strconv"
	"time"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type StageService struct{}

func (s *StageService) Sign(sign model.Sign, gra model.Grade) error {

	if err := global.DB.Where("user_id = ?", gra.UserId).First(&gra).Error; err != nil {
		// 如果没有找到，创建新用户
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := global.DB.Create(&gra).Error; err != nil {
				return err
			}
		}
	}
	// 报名表
	var count int64
	//global.DB.Model(&model.Sign{}).Where("user_id = ?", sign.UserId).Count(&count)
	subQuery := global.DB.Model(&model.User{}).Where("user_id = ?", sign.UserId).
		Joins("inner join user_authority on user_authority.user_id = user_id").
		Joins("inner join authorities on authority_id = user_authority.authority_authority_id").
		Where("authority_id = ?", 777).Select("1")

	global.DB.Model(&model.Sign{}).
		Where("user_id = ? AND jie_ci_id = ?", sign.UserId, sign.JieCiId).
		Where("NOT EXISTS (?)", subQuery).
		Count(&count)
	if count >= global.CONFIG.Sign.Number {
		errMsg := ",不能超出" + strconv.FormatInt(global.CONFIG.Sign.Number, 10) + "项报名内容"
		return errors.New(errMsg)
	}
	return global.DB.Create(&sign).Error
}

func (s *StageService) DeleteSign(id int) error {
	var sign model.Sign
	return global.DB.Where("id=?", id).Delete(&sign).Error
}

func (s *StageService) UpdateSign(signR model.Sign) error {
	return global.DB.Updates(&signR).Error
}

func (s *StageService) UpdateSignZero(data map[string]interface{}) error {
	return global.DB.Model(&model.Sign{}).Updates(&data).Error
}

func (s *StageService) UpdateSignCoverUrl(signID uint, coverUrl string) error {
	return global.DB.Model(&model.Sign{}).Where("id = ?", signID).Update("cover_url", coverUrl).Error
}

func (s *StageService) GetSign(signId int) (model.Sign, error) {
	var sign model.Sign
	err := global.DB.Where("id", signId).Preload("Files").First(&sign).Error
	return sign, err
}

func (s *StageService) GetSignEvaluate(signId int, userId uint) (model.Evaluate, error) {
	var evaluate model.Evaluate
	err := global.DB.Where("sign_id=? And evaluate_user_id=?", signId, userId).First(&evaluate).Error
	return evaluate, err
}

func (s *StageService) GetWorkFileType() ([]model.WorkFileType, error) {
	var workFileTypes []model.WorkFileType
	if err := global.DB.Find(&workFileTypes).Error; err != nil {
		return nil, err
	}
	return workFileTypes, nil
}

func (s *StageService) GetMajor() ([]model.Major, error) {
	var major []model.Major
	if err := global.DB.Find(&major).Error; err != nil {
		return nil, err
	}
	return major, nil
}

func (s *StageService) GetSignList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var signList []model.SignWithPhone
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
			switch fieldName {
			case "username":
				fieldName = "signs." + fieldName
				whereStr += fmt.Sprintf("%s = ? AND ", fieldName)
				whereArgs = append(whereArgs, val)
			case "major_id":
				// 从 grades 表中查找 major_id 的所有 userId
				var gradeUsers []model.Grade
				err = global.DB.Model(&model.Grade{}).Where("major_id = ?", val).Find(&gradeUsers).Error
				if err != nil {
					return
				}
				userIds := make([]uint, len(gradeUsers))
				for i, gradeUser := range gradeUsers {
					userIds[i] = gradeUser.UserId
				}
				// 在 signs 表中查找所有 userId 对应的 signs
				whereStr += fmt.Sprintf("user_id IN (?) AND ")
				whereArgs = append(whereArgs, userIds)
			default:
				fieldName = "signs." + fieldName
				if fieldName == "signs.status" {
					// 检查字段名是否为 status
					whereStr += fmt.Sprintf("%s IN (?) AND ", fieldName)
					whereArgs = append(whereArgs, val)
				} else {
					whereStr += fmt.Sprintf("%s = ? AND ", fieldName)
					whereArgs = append(whereArgs, val)
				}
			}
		}
		if len(whereArgs) > 0 {
			whereStr = whereStr[:len(whereStr)-4] // remove the last 'AND '
		}
		db = db.Debug().Where(whereStr, whereArgs...)
	}

	// 计数操作
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 执行查询
	err = db.Select("signs.*, users.phone as phone").Joins("left join users on signs.user_id = users.id").Limit(limit).Offset(offset).Scan(&signList).Error
	if err != nil {
		return
	}
	for i := range signList {
		var files []model.File
		err = global.DB.Model(&model.File{}).Where("sign_id = ?", signList[i].ID).Find(&files).Error
		if err != nil {
			return
		}
		signList[i].Files = files
	}
	return signList, total, err
}

func (s StageService) GetGrade(gradeName string) (uint, error) {
	grade := model.Grade{}
	err := global.DB.Where("grade_name = ?", gradeName).First(&grade).Error
	return grade.ID, err
}

func (s *StageService) GetJieCi() (model.CompetitionTime, error) {
	cpTime := model.CompetitionTime{}
	now := time.Now()
	err := global.DB.Where("start_time < ? AND end_time > ?", now, now).First(&cpTime).Error
	return cpTime, err
}

func (s *StageService) Upload(file model.File) (model.File, error) {
	return file, global.DB.Create(&file).Error
}

func (s *StageService) DeleteFile(fileId uint) error {
	keyWords := map[string]interface{}{
		"id": fileId,
	}
	fileFromDb, err := s.GetFile(keyWords)
	if err != nil {
		return err
	}
	if err = utils.DeleteFile(fileFromDb.Url); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.DB.Where("id = ?", fileId).Unscoped().Delete(&model.File{}).Error
	return err
}

func (s *StageService) GetFile(keyWords map[string]interface{}) (model.File, error) {
	var file model.File
	whereStr := ""
	whereArgs := []interface{}{}
	for key, val := range keyWords {
		fieldName := strcase.ToSnake(key) // 将小驼峰命名转换为下划线命名
		// 将数字类型的值转换为uint类型
		if v, ok := val.(float64); ok {
			val = uint(v)
		}
		whereStr += fmt.Sprintf("%s = ? ", fieldName)
		whereArgs = append(whereArgs, val)
		if len(whereArgs) != len(keyWords) {
			whereStr += "AND "
		}
	}
	err := global.DB.Where(whereStr, whereArgs...).First(&file).Error
	return file, err
}

func (s *StageService) GetFileList(keyWords map[string]interface{}) ([]model.File, error) {
	var fileList []model.File
	whereStr := ""
	whereArgs := []interface{}{}
	for key, val := range keyWords {
		fieldName := strcase.ToSnake(key) // 将小驼峰命名转换为下划线命名
		// 将数字类型的值转换为uint类型
		if v, ok := val.(float64); ok {
			val = uint(v)
		}
		whereStr += fmt.Sprintf("%s = ? ", fieldName)
		whereArgs = append(whereArgs, val)
		if len(whereArgs) != len(keyWords) {
			whereStr += "AND "
		}
	}
	err := global.DB.Where(whereStr, whereArgs...).Find(&fileList).Error
	return fileList, err
}

func (s *StageService) GetStage(currentTime time.Time) (model.CompetitionTime, error) {
	var cpTime model.CompetitionTime
	err := global.DB.Where("end_time > ? and start_time < ?", currentTime, currentTime).First(&cpTime).Error
	return cpTime, err
}

func (s *StageService) CheckSignExists(signId uint) (bool, error) {
	var count int64
	err := global.DB.Model(&model.Sign{}).Where("id = ?", signId).Count(&count).Error
	return count > 0, err
}

func (s *StageService) CheckLikeExists(userID, signId uint) (bool, error) {
	var count int64
	err := global.DB.Model(&model.UserLike{}).Where("user_id = ? AND sign_id = ?", userID, signId).Count(&count).Error
	return count > 0, err
}

func (s *StageService) CreateLike(userID, signId uint) error {
	like := model.UserLike{UserID: userID, SignID: signId}
	return global.DB.Create(&like).Error
}

func (s *StageService) IncrementLikes(signID int) error {
	return global.DB.Model(&model.Sign{}).Where("id = ?", signID).Update("likes", gorm.Expr("likes + 1")).Error
}
