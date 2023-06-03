package sys

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"strconv"
	"strings"
)

func AddMenu(rc *rctx.ReqCtx) {
	var menu form.MenusReq
	rctx.ShouldBind(rc, &menu)
	core.V1.Menu().Create(&menu)
}

func UpdateMenu(rc *rctx.ReqCtx) {
	var menu form.UpdateMenusReq
	rctx.ShouldBind(rc, &menu)
	menuID := rctx.PathParamUint64(rc, "id")
	core.V1.Menu().Update(&menu, menuID)
}

func DeleteMenu(rc *rctx.ReqCtx) {
	menuID := rctx.PathParamUint64(rc, "id")
	core.V1.Menu().Delete(menuID)
}

func GetMenu(rc *rctx.ReqCtx) {
	menuID := rctx.PathParamUint64(rc, "id")
	rc.ResData = core.V1.Menu().Get(menuID)
}

func ListMenus(rc *rctx.ReqCtx) {
	var menuType []int8
	isTree := true

	menuTypeStr := rctx.QueryDefault(rc, "menuType", "1,2,3")
	tree := rctx.QueryDefault(rc, "isTree", "true")

	if tree == "false" {
		isTree = false
	}

	menuTypeSlice := strings.Split(menuTypeStr, ",")
	for _, t := range menuTypeSlice {
		res, err := strconv.Atoi(t)
		restfulx.ErrIsNil(err, restfulx.ParamErr)
		menuType = append(menuType, int8(res))
	}
	page := rctx.QueryDefaultInt(rc, "page", 0)
	limit := rctx.QueryDefaultInt(rc, "limit", 0)

	rc.ResData = core.V1.Menu().List(page, limit, menuType, isTree)
}

func UpdateMenuStatus(rc *rctx.ReqCtx) {
	menuID := rctx.PathParamUint64(rc, "id")
	status := rctx.PathParamUint64(rc, "status")
	core.V1.Menu().UpdateStatus(menuID, status)
}
