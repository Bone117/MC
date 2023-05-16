package service

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/common/request"

	"gorm.io/gorm"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct{}

//var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(auth model.Authority) (authority model.Authority, err error) {
	var authorityBox model.Authority
	if !errors.Is(global.DB.Where("authority_id=?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = global.DB.Create(&auth).Error
	return auth, err
}

func (authorityService *AuthorityService) UpdateAuthority(auth model.Authority) (authority model.Authority, err error) {
	err = global.DB.Where("authority_id = ?", auth.AuthorityId).First(&model.Authority{}).Updates(&auth).Error
	return auth, err
}

func (authorityService *AuthorityService) DeleteAuthority(auth *model.Authority) (err error) {
	if errors.Is(global.DB.Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	err = global.DB.Delete(&[]model.UseAuthority{}, "authority_authority_id=?", auth.AuthorityId).Error
	ServiceGroupApp.CasbinService.ClearCasbin(0, auth.AuthorityId)
	return err
}

func (authorityService *AuthorityService) GetAuthorityList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	var authorityList []model.Authority
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.DB.Model(&model.Authority{})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Users").Find(&authorityList).Error
	return authorityList, total, err
}
