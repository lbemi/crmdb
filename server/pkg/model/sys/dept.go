package sys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"gorm.io/gorm"
	"time"
)

type Department struct {
	model.Model
	DeptName string `json:"deptName" gorm:"column:deptName"`
	Remark   string `json:"remark" gorm:"column:remark"`
	Status   int8   `json:"status" gorm:"column:status"`
	Sort     int8   `json:"sort" gorm:"sort"`
}

// TableName 自定义表名
func (d *Department) TableName() string {
	return "departments"
}

// BeforeCreate 创建前操作
func (d *Department) BeforeCreate(*gorm.DB) error {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate 更新前操作
func (d *Department) BeforeUpdate(*gorm.DB) error {
	d.UpdatedAt = time.Now()
	return nil
}
