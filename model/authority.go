package model

import "time"

type Authority struct {
	AuthorityId   string     `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName string     `json:"authorityName" gorm:"comment:角色名"`
	Users         []User     `json:"-" gorm:"many2many:user_authority;"`
	CreatedAt     time.Time  // 创建时间
	UpdatedAt     time.Time  // 更新时间
	DeletedAt     *time.Time `sql:"index"`
}
