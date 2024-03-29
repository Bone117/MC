package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint           `gorm:"primarykey"`             // 主键ID
	CreatedAt time.Time      `gorm:"index" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"index" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`         // 删除时间
}
