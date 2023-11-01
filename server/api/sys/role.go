package sys

import (
	"strconv"
	"strings"

	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
)

func AddRole(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var role form.RoleReq
	rc.ShouldBind(&role)
	core.V1.Role().Create(c, &role)
}

func UpdateRole(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var role form.UpdateRoleReq
	rc.ShouldBind(&role)
	roleID := rc.PathParamUint64("id")
	core.V1.Role().Update(c, &role, roleID)
}

func DeleteRole(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	rid := rc.PathParamUint64("id")
	core.V1.Role().Delete(c, rid)
}

func GetRole(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	rid := rc.PathParamUint64("id")
	rc.ResData = core.V1.Role().Get(c, rid)
}

func ListRoles(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	condition := &sys.Role{}
	query := &model.PageParam{}
	query.Page = rc.QueryDefaultInt("page", 0)
	query.Limit = rc.QueryDefaultInt("limit", 0)
	condition.Name = rc.Query("name")
	condition.Status = rc.QueryParamInt8("status")
	rc.ResData = core.V1.Role().List(c, query, condition)
}

func GetMenusByRole(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
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
	rc.ResData = core.V1.Role().GetMenusByRoleID(c, rid, menuType)
}

func SetRoleMenus(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	rid := rc.PathParamUint64("id")
	var menuIDs form.Menus
	rc.ShouldBind(&menuIDs)
	core.V1.Role().SetRole(c, rid, menuIDs.MenuIDS)
}

func UpdateRoleStatus(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	status := rc.PathParamUint64("status")
	roleID := rc.PathParamUint64("id")
	core.V1.Role().UpdateStatus(c, roleID, status)
}
