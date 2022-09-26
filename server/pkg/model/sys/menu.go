package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

// Menu 菜单
type Menu struct {
	basemodel.Model
	Status   int8   `gorm:"column:status;type:tinyint(1);not null;comment:状态(1:启用 2:不启用)" json:"status" form:"status"`                               // 状态(1:启用 2:不启用)
	Memo     string `gorm:"column:memo;size:128;comment:备注" json:"memo" form:"memo"`                                                                 // 备注
	ParentID int64  `gorm:"column:parent_id;not null;comment:父级ID" json:"parent_id" form:"parent_id"`                                                // 父级ID
	URL      string `gorm:"column:url;size:128;comment:菜单URL" json:"url,omitempty" form:"url"`                                                       // 菜单URL
	Name     string `gorm:"column:name;size:128;not null;comment:菜单名称" json:"name" form:"name"`                                                      // 菜单名称
	Sequence int    `gorm:"column:sequence;not null;comment:排序值" json:"sequence" form:"sequence"`                                                    // 排序值
	MenuType int8   `gorm:"column:menu_type;type:tinyint(1);not null;comment:菜单类型(1 左侧菜单,2 按钮, 3 非展示权限)" json:"menu_type" form:"menu_type"`          // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Icon     string `gorm:"column:icon;size:32;comment:icon图标" json:"icon" form:"icon"`                                                              // icon
	Method   string `gorm:"column:method;size:32;not null;default:none;comment:操作类型 none/GET/POST/PUT/DELETE" json:"method,omitempty" form:"method"` // 操作类型 none/GET/POST/PUT/DELETE
	Children []Menu `gorm:"-" json:"children"`
}

// TableName 表名
func (m *Menu) TableName() string {
	return "menus"
}

// BeforeCreate 添加前
func (m *Menu) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前
func (m *Menu) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
