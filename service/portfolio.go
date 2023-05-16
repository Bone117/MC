package service

import (
	"fmt"
	"server/global"
	"server/model"
	"server/model/common/request"
)

type PortfolioService struct{}

func (p *PortfolioService) Upload(file model.Portfolio) (model.Portfolio, error) {
	return file, global.DB.Create(&file).Error
}

func (p *PortfolioService) GetPastWork(Id int) (model.Portfolio, error) {
	var work model.Portfolio
	err := global.DB.Where("id", Id).First(&work).Error
	return work, err
}

func (p *PortfolioService) GetPortfolioList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var portfolioList []model.Portfolio
	limit := pageInfo.PageSize
	offset := pageInfo.Page
	db := global.DB.Model(&model.Portfolio{})
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
		err = db.Limit(limit).Offset(offset).Order("created_at DESC").Where(whereStr, whereArgs...).Find(&portfolioList).Statement.Error
	} else {
		err = db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&portfolioList).Error
	}

	return portfolioList, total, err
}
