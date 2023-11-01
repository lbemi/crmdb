package sys

import (
	"strconv"
	"strings"

	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
)

func AddMenu(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var menu form.MenusReq
	rc.ShouldBind(&menu)
	core.V1.Menu().Create(c, &menu)
}

func UpdateMenu(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var menu form.UpdateMenusReq
	rc.ShouldBind(&menu)
	menuID := rc.PathParamUint64("id")
	core.V1.Menu().Update(c, &menu, menuID)
}

func DeleteMenu(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	menuID := rc.PathParamUint64("id")
	core.V1.Menu().Delete(c, menuID)
}

func GetMenu(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	menuID := rc.PathParamUint64("id")
	rc.ResData = core.V1.Menu().Get(c, menuID)
}

func ListMenus(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var menuType []int8
	condition := &sys.Menu{}
	isTree := true

	menuTypeStr := rc.QueryDefault("menuType", "1,2,3")
	tree := rc.QueryDefault("isTree", "true")
	condition.Group = rc.Query("group")

	if tree == "false" {
		isTree = false
	}

	menuTypeSlice := strings.Split(menuTypeStr, ",")
	for _, t := range menuTypeSlice {
		res, err := strconv.Atoi(t)
		restfulx.ErrNotNilDebug(err, restfulx.ParamErr)
		menuType = append(menuType, int8(res))
	}
	page := rc.QueryDefaultInt("page", 0)
	limit := rc.QueryDefaultInt("limit", 0)
	condition.Memo = rc.Query("memo")
	condition.Status = rc.QueryParamInt8("status")
	rc.ResData = core.V1.Menu().List(c, page, limit, menuType, isTree, condition)
}

func UpdateMenuStatus(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	menuID := rc.PathParamUint64("id")
	status := rc.PathParamUint64("status")
	core.V1.Menu().UpdateStatus(c, menuID, status)
}
