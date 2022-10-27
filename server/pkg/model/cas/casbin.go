package cas

type CasbinModel struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	PType  string `json:"ptype" gorm:"column:ptype;size:100" description:"策略类型"`
	Role   string `json:"role" gorm:"column:v0;size:100" description:"角色"`
	Path   string `json:"path" gorm:"column:v1;size:100" description:"api路径"`
	Method string `json:"method" gorm:"column:v2;size:100" description:"访问方法"`
	//V3     string `gorm:"size:192;uniqueIndex:unique_index"`
	//V4     string `gorm:"size:192;uniqueIndex:unique_index"`
	//V5     string `gorm:"size:192;uniqueIndex:unique_index"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}
