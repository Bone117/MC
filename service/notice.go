package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
)

type NoticeService struct{}

func (n *NoticeService) CreateNotice(notice model.Notice) (uint, error) {
	err := global.DB.Create(&notice).Error
	return notice.ID, err
}

func (n *NoticeService) DeleteNotice(noticeID uint) error {
	err := global.DB.Where("id=?", noticeID).Delete(&model.Notice{}).Error
	return err
}

func (n *NoticeService) UpdateNotice(notice model.Notice) error {
	return global.DB.Updates(&notice).Error
}

func (n *NoticeService) GetNotice(noticeId uint) (model.Notice, error) {
	var notice model.Notice
	err := global.DB.Where("id", noticeId).First(&notice).Error
	return notice, err
}

func (n NoticeService) GetNoticeList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var noticeList []model.Notice
	limit := pageInfo.PageSize
	//offset := pageInfo.PageSize * (pageInfo.Page - 1)
	offset := pageInfo.Page
	db := global.DB.Model(&model.Notice{})
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
		err = db.Limit(limit).Offset(offset).Order("created_at DESC").Where(whereStr, whereArgs...).Find(&noticeList).Statement.Error
	} else {
		err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&noticeList).Error
	}

	return noticeList, total, err
}
