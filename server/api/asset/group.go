package asset

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func AddGroup(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	group := &asset.Group{}
	rc.ShouldBind(group)
	core.V1.Group().Create(c, group)
}

func ListGroup(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	page := rc.GetPageQueryParam()
	rc.ResData = core.V1.Group().List(c, page.Page, page.Limit)
}

func GetGroup(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Group().GetByGroupId(c, id)
}

func DeleteGroup(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	core.V1.Group().Delete(c, id)
}

func UpdateGroup(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	group := &asset.Group{}
	rc.ShouldBind(group)
	core.V1.Group().Update(c, group.ID, group)
}
