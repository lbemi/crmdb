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
	rc.ShouldBind(&role)
	core.V1.Role().Create(&role)
}

func UpdateRole(rc *rctx.ReqCtx) {
	var role form.UpdateRoleReq
	rc.ShouldBind(&role)
	roleID := rc.PathParamUint64("id")
	core.V1.Role().Update(&role, roleID)
}

func DeleteRole(rc *rctx.ReqCtx) {
	rid := rc.PathParamUint64("id")
	core.V1.Role().Delete(rid)
}

func GetRole(rc *rctx.ReqCtx) {
	rid := rc.PathParamUint64("id")
	rc.ResData = core.V1.Role().Get(rid)
}

func ListRoles(rc *rctx.ReqCtx) {
	condition := &sys.Role{}
	query := &model.PageParam{}
	query.Page = rc.QueryDefaultInt("page", 0)
	query.Limit = rc.QueryDefaultInt("limit", 0)
	condition.Name = rc.QueryParam("name")
	condition.Status = rc.QueryParamInt8("status")
	rc.ResData = core.V1.Role().List(query, condition)
}

func GetMenusByRole(rc *rctx.ReqCtx) {
	rid := rc.PathParamUint64("id")
	menuTypeStr := rc.QueryDefault("menuType", "1,2,3")
	var menuType []int8
	menuTypeSlice := strings.Split(menuTypeStr, ",")
	for _, t := range menuTypeSlice {
		res, err := strconv.Atoi(t)
		if err != nil {
			restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
			return
		}
		menuType = append(menuType, int8(res))
	}
	rc.ResData = core.V1.Role().GetMenusByRoleID(rid, menuType)
}

func SetRoleMenus(rc *rctx.ReqCtx) {
	rid := rc.PathParamUint64("id")
	var menuIDs form.Menus
	rc.ShouldBind(&menuIDs)
	core.V1.Role().SetRole(rid, menuIDs.MenuIDS)
}

func UpdateRoleStatus(rc *rctx.ReqCtx) {
	status := rc.PathParamUint64("status")
	roleID := rc.PathParamUint64("id")
	core.V1.Role().UpdateStatus(roleID, status)
}
