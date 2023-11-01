package form

import (
	"github.com/lbemi/lbemi/apps/asset/entity"
)

// PageHost 分页菜单
type PageHost struct {
	Hosts []entity.Host `json:"hosts"`
	Total int64         `json:"total"`
}
