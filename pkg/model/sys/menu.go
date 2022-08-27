package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

//Menu 菜单
type Menu struct {
	basemodel.Model
	Status      uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`             // 状态(1:启用 2:不启用)
	Memo        string `gorm:"column:memo;size:64;" json:"memo" form:"memo"`                                    // 备注
	ParentID    uint64 `gorm:"column:parent_id;not null;" json:"parent_id" form:"parent_id"`                    // 父级ID
	URL         string `gorm:"column:url;size:72;" json:"url" form:"url"`                                       // 菜单URL
	Name        string `gorm:"column:name;size:32;not null;" json:"name" form:"name"`                           // 菜单名称
	Sequence    int    `gorm:"column:sequence;not null;" json:"sequence" form:"sequence"`                       // 排序值
	MenuType    uint8  `gorm:"column:menu_type;type:tinyint(1);not null;" json:"menu_type" form:"menu_type"`    // 菜单类型 1 模块,2菜单,3操作
	Code        string `gorm:"column:code;size:32;not null;unique_index:uk_menu_code;" json:"code" form:"code"` // 菜单代码
	Icon        string `gorm:"column:icon;size:32;" json:"icon" form:"icon"`                                    // icon
	OperateType string `gorm:"column:operate_type;size:32;not null;" json:"operate_type" form:"operate_type"`   // 操作类型 none/add/del/view/update
}

//TableName 表名
func (Menu) TableName() string {
	return "menu"
}

//BeforeCreate 添加前
func (m *Menu) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *Menu) BeforeUpdate(*gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
