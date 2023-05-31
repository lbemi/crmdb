package sys

import (
	"time"

	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/model/basemodel"
)

// Menu 菜单
type Menu struct {
	basemodel.Model
	Status    int8   `gorm:"column:status;type:tinyint(1);not null;default:1;comment:状态(1:启用 2:不启用)" json:"status" form:"status"`          // 状态(1:启用 2:不启用)
	Memo      string `gorm:"column:memo;size:128;comment:备注" json:"memo" form:"memo"`                                                      // 备注
	ParentID  uint64 `gorm:"column:parentID;not null;comment:父级ID" json:"parentID" form:"parent_id"`                                       // 父级ID
	Path      string `gorm:"column:path;size:128;comment:菜单URL" json:"path,omitempty" form:"path"`                                         // 菜单URL
	Component string `gorm:"column:component;" json:"component"`                                                                           // 前端组件
	Name      string `gorm:"column:name;size:128;not null;comment:菜单名称" json:"name" form:"name"`                                           // 菜单名称
	Sequence  int    `gorm:"column:sequence;not null;comment:排序值" json:"sequence" form:"sequence"`                                         // 排序值
	MenuType  int8   `gorm:"column:menuType;type:tinyint(1);not null;comment:菜单类型(1 左侧菜单,2 按钮, 3 非展示权限)" json:"menuType" form:"menu_type"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Redirect  string `json:"redirect" gorm:"column:redirect;size:128;"`
	Method    string `gorm:"column:method;size:32;not null;comment:操作类型 none/GET/POST/PUT/DELETE" json:"method,omitempty" form:"method"` // 操作类型 none/GET/POST/PUT/DELETE
	Code      string `gorm:"column:code;size:128;not null;" json:"code"`                                                                 // 前端鉴权code 例： user:role:add, user:role:delete
	Group     string `gorm:"column:group;size:256;not null;" json:"group"`
	Meta      `json:"meta"`
	Children  []Menu `gorm:"-" json:"children"`
}

type Meta struct {
	Title       string `json:"title" gorm:"column:title;comment:标题"`
	IsLink      string `json:"isLink" gorm:"column:isLink;type:varchar(256)"`
	IsHide      bool   `json:"isHide" gorm:"column:isHide"`
	IsKeepAlive bool   `json:"isKeepAlive" gorm:"column:isKeepAlive"`
	IsAffix     bool   `json:"isAffix" gorm:"column:isAffix"`
	IsIframe    bool   `json:"isIframe" gorm:"column:isIframe"`
	IsK8s       bool   `json:"isK8S" gorm:"column:isK8S"`
	Icon        string `gorm:"column:icon;size:32;comment:icon图标" json:"icon" form:"icon"`
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
