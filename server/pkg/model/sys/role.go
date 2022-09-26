package sys

import (
	"time"

	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"github.com/lbemi/lbemi/pkg/util"
)

// Role 角色
type Role struct {
	basemodel.Model
	Memo     string `gorm:"column:memo;size:128;comment:备注" json:"memo" form:"memo"`                                    // 备注
	Name     string `gorm:"column:name;size:128;not null;unique_index:uk_role_name;comment:名称" json:"name" form:"name"` // 名称
	Sequence int    `gorm:"column:sequence;not null;comment:排序值" json:"sequence" form:"sequence"`                       //
	ParentID int64  `gorm:"column:parent_id;not null;comment:父级ID" json:"parent_id" form:"parent_id"`                   // 父级ID
	Status   int8   `gorm:"column:status;not null;default:1;comment:状态：0 表示禁用，1 表示启用" json:"status" form:"status" `     // 0 表示禁用，1 表示启用
	Children []Role `gorm:"-"`                                                                                          // 子角色信息
}

// TableName 自定义表名
func (r *Role) TableName() string {
	return "roles"
}

// BeforeCreate 创建前操作
func (r *Role) BeforeCreate(*gorm.DB) error {
	r.ID = util.GetSnowID()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前操作
func (r *Role) BeforeUpdate(*gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Role) GetRoleByUId(uid string) (err error) {
	if err = global.App.DB.Where("id = ?", uid).First(r).Error; err != nil {
		return
	}
	return
}
