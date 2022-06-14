package service

import (
	"errors"
	"server/global"
	"server/model"

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
