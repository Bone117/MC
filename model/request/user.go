package request

import "server/model"

type Register struct {
	Username string `json:"username"`
	NickName string `json:"nickName" gorm:"default:'newUser'"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	//AuthorityId  string   `json:"authorityId"`
	AuthorityIds []string `json:"authorityIds"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Captcha   string `json:"captcha"`
	//CaptchaId string `json:"captchaId"` // 验证码ID
}

type ChangePasswordStruct struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type ChangeUserInfo struct {
	ID           uint              `gorm:"primarykey"`
	NickName     string            `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Phone        string            `json:"phone"  gorm:"comment:用户手机号"`               // 用户角色ID
	AuthorityIds []string          `json:"authorityIds" gorm:"-"`                     // 角色ID
	Email        string            `json:"email"  gorm:"comment:用户邮箱"`                // 用户邮箱
	Authorities  []model.Authority `json:"-" gorm:"many2many:user_authority;"`
}

type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}
