package model

type UserAuthority struct {
	UserId               uint   `gorm:"column:user_id"`
	AuthorityAuthorityId string `gorm:"column:authority_authority_id"`
}
