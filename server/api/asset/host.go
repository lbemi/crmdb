package asset

import (
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
)

func AddHost(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var machine asset.Host
	rc.ShouldBind(&machine)
	core.V1.Host().Create(c, &machine)
}

func ListHosts(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	pageParam := rc.GetPageQueryParam()
	rc.ResData = core.V1.Host().List(c, pageParam.Page, pageParam.Limit)
}

func GetHostById(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Host().GetByHostId(c, id)
}

func UpdateHost(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var machine asset.Host
	rc.ShouldBind(&machine)
	core.V1.Host().Update(c, machine.ID, &machine)
}

func DeleteHost(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	core.V1.Host().Delete(c, id)
}

func WsShell(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	account_id := rc.PathParamUint64("account_id")
	col := rc.QueryDefaultInt("cols", 170)
	row := rc.QueryDefaultInt("rows", 38)
	client, session, channel := core.V1.Terminal().GenerateClient(c, id, account_id, col, row)
	conn, err := wsstore.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	core.V1.Ws().GenerateConn(conn, client, session, channel)
}
