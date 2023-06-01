package sys

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
	"strings"
)

func AddRole(rc *rctx.ReqCtx) {
	var role form.RoleReq
	rctx.ShouldBind(rc, &role)
	core.V1.Role().Create(&role)
}

func UpdateRole(rc *rctx.ReqCtx) {
	var role form.UpdateRoleReq
	rctx.ShouldBind(rc, &role)
	roleID := rctx.ParamUint64(rc, "id")
	core.V1.Role().Update(&role, roleID)
}

func DeleteRole(rc *rctx.ReqCtx) {
	rid := rctx.ParamUint64(rc, "id")
	core.V1.Role().Delete(rid)
}

func GetRole(rc *rctx.ReqCtx) {
	rid := rctx.ParamUint64(rc, "id")
	rc.ResData = core.V1.Role().Get(rid)
}

func ListRoles(rc *rctx.ReqCtx) {
	condition := &sys.Role{}
	query := &model.PageParam{}
	query.Page = rctx.QueryDefaultInt(rc, "page", 0)
	query.Limit = rctx.QueryDefaultInt(rc, "limit", 0)
	condition.Name = rctx.QueryParam(rc, "name")
	rc.ResData = core.V1.Role().List(query, condition)
}

func GetMenusByRole(rc *rctx.ReqCtx) {
	rid := rctx.ParamUint64(rc, "id")
	menuTypeStr := rctx.QueryDefault(rc, "menuType", "1,2,3")
	var menuType []int8
	menuTypeSlice := strings.Split(menuTypeStr, ",")
	for _, t := range menuTypeSlice {
		res, err := strconv.Atoi(t)
		if err != nil {
			restfulx.ErrIsNilRes(err, restfulx.ParamErr)
			return
		}
		menuType = append(menuType, int8(res))
	}
	rc.ResData = core.V1.Role().GetMenusByRoleID(rid, menuType)
}

func SetRoleMenus(rc *rctx.ReqCtx) {
	rid := rctx.ParamUint64(rc, "id")
	var menuIDs form.Menus
	rctx.ShouldBind(rc, &menuIDs)
	core.V1.Role().SetRole(rid, menuIDs.MenuIDS)
}

func UpdateRoleStatus(rc *rctx.ReqCtx) {
	status := rctx.ParamUint64(rc, "status")
	roleID := rctx.ParamUint64(rc, "id")
	core.V1.Role().UpdateStatus(roleID, status)
}
