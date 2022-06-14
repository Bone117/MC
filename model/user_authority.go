package model

type UseAuthority struct {
	UserId               uint   `gorm:"column:user_id"`
	AuthorityAuthorityId string `gorm:"column:authority_authority_id"`
}

func (s *UseAuthority) TableName() string {
	return "user_authority"
}
