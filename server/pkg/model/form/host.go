package form

import (
	"github.com/lbemi/lbemi/pkg/model/asset"
)

// PageHost 分页菜单
type PageHost struct {
	Hosts []asset.Host `json:"hosts"`
	Total int64        `json:"total"`
}
