package sys

import "github.com/lbemi/lbemi/pkg/model/basemodel"

type Dept struct {
	basemodel.Model
	DeptName string `json:"deptName" gorm:"column:deptName"`
	Remark   string `json:"remark" gorm:"column:remark"`
	Status   int8   `json:"status" gorm:"column:status"`
	Sort     int8   `json:"sort" gorm:"sort"`
}
