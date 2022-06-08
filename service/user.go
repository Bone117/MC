package service

import (
	"errors"
	"fmt"
	"server/global"
	"server/model"
	"server/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

// Register 用户注册
func (userService *UserService) Register(u model.User) (userInter model.User, err error) {
	var user model.User
	// 判断是否注册
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.DB.Create(&u).Error
	return u, err
}

// Login 用户登录
func (userService *UserService) Login(u *model.User) (err error, userInter *model.User) {
	if nil == global.DB {
		return fmt.Errorf("db not init"), nil
	}

	var user model.User
	err = global.DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return errors.New("密码错误"), nil
		}
	}
	return nil, &model.User{}
}
