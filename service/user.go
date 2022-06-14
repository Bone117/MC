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
func (userService *UserService) Login(u *model.User) (userInter *model.User, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user model.User
	err = global.DB.Where("username = ?", u.Username).Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, nil
}

// ChangePassword 修改密码
func (userService *UserService) ChangePassword(u *model.User, newPassword string) (userInter *model.User, err error) {
	var user model.User
	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.DB.Save(&user).Error
	return &user, err
}

// ResetPassword 重置密码
func (userService *UserService) ResetPassword(id uint) (err error) {
	err = global.DB.Model(&model.User{}).Where("id = ?", id).Update("password", utils.BcryptHash("123456")).Error
	return err
}

// SetUserInfo 设置用户信息
func (userService *UserService) SetUserInfo(req model.User) error {
	return global.DB.Updates(&req).Error
}

// DeleteUser 删除用户
func (userService *UserService) DeleteUser(id int) (err error) {
	var user model.User
	err = global.DB.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = global.DB.Delete(&[]model.UseAuthority{}, "user_id=?", id).Error
	return err
}

// TODO 未完成
func (userService *UserService) setUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	assignErr := global.DB.Where("user_id = ? AND authority_authority_id = ?", id, authorityId).First(&model.UseAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.DB.Where("uuid = ?", uuid).First(&model.User{}).Update("authority_id", authorityId).Error
	return err
}

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]model.UseAuthority{}, "user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []model.UseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, model.UseAuthority{
				id, v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&model.User{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}
