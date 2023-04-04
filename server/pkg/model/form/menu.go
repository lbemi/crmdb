package form

import "github.com/lbemi/lbemi/pkg/model/sys"

type Menus struct {
	MenuIDS []uint64 `json:"menuIDS"`
}

type MenusReq struct {
	Status    int8   `json:"status"`   // 状态(1:启用 2:不启用)
	Memo      string `json:"memo"`     // 备注
	ParentID  uint64 `json:"parentID"` // 父级ID
	Path      string `json:"path"`     // 菜单URL
	Name      string `json:"name"`     // 菜单名称
	Sequence  int    `json:"sequence"` // 排序值
	MenuType  int8   `json:"menuType"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Method    string `json:"method"`   // 操作类型 none/GET/POST/PUT/DELETE
	Code      string `json:"code"`     // 前端鉴权code 例： user:role:add, user:role:delete
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Group     string `json:"group"`
	Meta      `json:"meta"`
}
type Meta struct {
	Title       string `json:"title"`
	IsLink      string `json:"isLink"`
	IsHide      bool   `json:"isHide"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
	IsK8s       bool   `json:"isK8S" gorm:"isK8s"`
	Icon        string `json:"icon"`
}

type UpdateMenusReq struct {
	Status    int8   `json:"status"`   // 状态(1:启用 2:不启用)
	Memo      string `json:"memo"`     // 备注
	ParentID  uint64 `json:"parentID"` // 父级ID
	Path      string `json:"path"`     // 菜单URL
	Name      string `json:"name"`     // 菜单名称
	Sequence  int    `json:"sequence"` // 排序值
	MenuType  int8   `json:"menuType"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Icon      string `json:"icon"`     // icon
	Method    string `json:"method"`   // 操作类型 none/GET/POST/PUT/DELETE
	Code      string `json:"code"`     // 前端鉴权code 例： user:role:add, user:role:delete
	Component string `json:"component"`
	Group     string `json:"group"`
	Redirect  string `json:"redirect"`
	Meta      `json:"meta" gorm:"meta"`
}

// PageMenu 分页菜单
type PageMenu struct {
	Menus []sys.Menu `json:"menus"`
	Total int64      `json:"total"`
}
type Roles struct {
	RoleIds []uint64 `json:"role_ids"`
}
