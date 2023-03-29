package form

import "github.com/lbemi/lbemi/pkg/model/sys"

type Menus struct {
	MenuIDS []uint64 `json:"menu_ids"`
}

type MenusReq struct {
	Status   int8   `json:"status"`   // 状态(1:启用 2:不启用)
	Memo     string `json:"memo"`     // 备注
	ParentID uint64 `json:"parentID"` // 父级ID
	Path     string `json:"path"`     // 菜单URL
	Name     string `json:"name"`     // 菜单名称
	Sequence int    `json:"sequence"` // 排序值
	MenuType int8   `json:"menuType"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Icon     string `json:"icon"`     // icon
	Method   string `json:"method"`   // 操作类型 none/GET/POST/PUT/DELETE
	Code     string `json:"code"`     // 前端鉴权code 例： user:role:add, user:role:delete
	Meta     `json:"meta"`
}
type Meta struct {
	Title       string `json:"title"`
	IsLink      bool   `json:"isLink"`
	IsHide      bool   `json:"isHide"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
	Icon        string `json:"icon"`
}

type UpdateMenusReq struct {
	Status   int8   `json:"status"`    // 状态(1:启用 2:不启用)
	Memo     string `json:"memo"`      // 备注
	ParentID uint64 `json:"parent_id"` // 父级ID
	Path     string `json:"path"`      // 菜单URL
	Name     string `json:"name"`      // 菜单名称
	Sequence int    `json:"sequence"`  // 排序值
	MenuType int8   `json:"menu_type"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Icon     string `json:"icon"`      // icon
	Method   string `json:"method"`    // 操作类型 none/GET/POST/PUT/DELETE
	Code     string `json:"code"`      // 前端鉴权code 例： user:role:add, user:role:delete
}

// PageMenu 分页菜单
type PageMenu struct {
	Menus []sys.Menu `json:"menus"`
	Total int64      `json:"total"`
}
type Roles struct {
	RoleIds []uint64 `json:"role_ids"`
}
