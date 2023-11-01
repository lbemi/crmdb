package config

type Rule struct {
	ID     uint64 `gorm:"column:id;primary_key;" json:"id" form:"id"`
	PType  string `json:"ptype" gorm:"column:ptype;size:100" description:"策略类型"`
	Role   string `json:"role" gorm:"column:v0;size:100" description:"角色"`
	Path   string `json:"path" gorm:"column:v1;size:100" description:"api路径"`
	Method string `json:"method" gorm:"column:v2;size:100" description:"访问方法"`
	V3     string `gorm:"column:v3;size:100"`
	V4     string `gorm:"column:v4;size:100"`
	V5     string `gorm:"column:v5;size:100"`
}

func (r *Rule) TableName() string {
	return "rules"
}
