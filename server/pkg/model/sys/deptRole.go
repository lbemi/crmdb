package sys

import (
	"github.com/lbemi/lbemi/pkg/model/basemodel"
	"gorm.io/gorm"
	"time"
)

type DepartmentRole struct {
	basemodel.Model
	DeptName string `json:"deptName" gorm:"column:deptName"`
	Remark   string `json:"remark" gorm:"column:remark"`
	Status   int8   `json:"status" gorm:"column:status"`
	Sort     int8   `json:"sort" gorm:"sort"`
}

// TableName 自定义表名
func (d *DepartmentRole) TableName() string {
	return "deptRoles"
}

// BeforeCreate 创建前操作
func (d *DepartmentRole) BeforeCreate(*gorm.DB) error {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前操作
func (d *DepartmentRole) BeforeUpdate(*gorm.DB) error {
	d.UpdatedAt = time.Now()
	return nil
}
