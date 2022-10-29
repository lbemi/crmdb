package basemodel

import (
	"time"
)

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT;" json:"id" form:"id"`                     // 主键
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at" form:"updated_at"` // 更新时间
}
