package form

import "github.com/lbemi/lbemi/pkg/model/sys"

type PageRole struct {
	Roles []sys.Role `json:"roles"`
	Total int64      `json:"total"`
}

type RoleReq struct {
	Memo     string `json:"memo" `      // 备注
	Name     string `json:"name"`       // 名称
	Sequence int    `json:"sequence" `  // 排序值
	ParentID uint64 `json:"parent_id" ` // 父级ID
	Status   int8   `json:"status" `    // 2 表示禁用，1 表示启用
}

type UpdateRoleReq struct {
	Memo     string `json:"memo" `      // 备注
	Name     string `json:"name"`       // 名称
	Sequence int    `json:"sequence" `  // 排序值
	ParentID uint64 `json:"parent_id" ` // 父级ID
	Status   int8   `json:"status" `    // 0 表示禁用，1 表示启用
}
