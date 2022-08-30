package cas

import (
	"errors"
	"github.com/lbemi/lbemi/pkg/global"
	"gorm.io/gorm"
)

type CasbinModel struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	PType  string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	Role   string `json:"role" gorm:"column:v0" description:"角色"`
	Path   string `json:"path" gorm:"column:v1" description:"api路径"`
	Method string `json:"method" gorm:"column:v2" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Create(db *gorm.DB) error {
	e := global.App.Enforcer
	success, err := e.AddPolicy(c.Role, c.Path, c.Method)
	if success == false {
		return errors.New("存在相同的API，添加失败")
	}
	if err != nil {
		return err
	}
	return nil
}

//func (c *CasbinModel) Update(db *gorm.DB, values interface{}) error {
//	if err := db.Model(c).Where("v1 = ? AND v2 = ?", c.Path, c.Method).Update(values).Error; err != nil {
//		return err
//	}
//	return nil
//}

func (c *CasbinModel) List(db *gorm.DB) [][]string {
	e := global.App.Enforcer
	policy := e.GetFilteredPolicy(0, c.Role)
	return policy
}
