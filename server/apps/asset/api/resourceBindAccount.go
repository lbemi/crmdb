package api

import (
	"github.com/lbemi/lbemi/apps/asset/entity"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func BindAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	ha := &entity.HostAccount{}
	rc.ShouldBind(ha)
	core.V1.ResourceBindAccount().BindAccount(c, ha)
}

func List(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	page := rc.GetPageQueryParam()
	rc.ResData = core.V1.ResourceBindAccount().List(c, page.Page, page.Limit)
}

func Get(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.ResourceBindAccount().Get(c, id)
}

func UnbindAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	core.V1.ResourceBindAccount().UnbindAccount(c, id)
}

func UpdateHostAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	ha := &entity.HostAccount{}
	rc.ShouldBind(ha)
	core.V1.ResourceBindAccount().UpdateHostAccount(c, ha)
}
