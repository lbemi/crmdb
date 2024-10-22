package api

import (
	"github.com/lbemi/lbemi/apps/asset/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
)

func AddHost(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var machine entity.Host
	rc.ShouldBind(&machine)
	core.V1.Host().Create(c, &machine)
}

func ListHosts(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	pageParam := rc.GetPageQueryParam()
	groupStr := rc.Query("groups")
	groups := util.ConvertSliceStrToInt(groupStr)
	ip := rc.Query("ip")
	label := rc.Query("label")
	description := rc.Query("description")
	rc.ResData = core.V1.Host().List(c, pageParam.Page, pageParam.Limit, groups, ip, label, description)
}

func GetHostById(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Host().GetByHostId(c, id)
}
func GetHostByGroups(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	pageParam := rc.GetPageQueryParam()
	group := entity2.GroupVo{}
	rc.ShouldBind(&group)
	rc.ResData = core.V1.Host().GetByGroup(c, group.Groups, pageParam.Page, pageParam.Limit)
}

func UpdateHost(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var machine entity.Host
	rc.ShouldBind(&machine)
	core.V1.Host().Update(c, machine.ID, &machine)
}

func DeleteHost(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	core.V1.Host().Delete(c, id)
}

func GetHostAccounts(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Host().GetHostAccounts(c, id)
}

func WsShell(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	accountId := rc.PathParamUint64("account")
	col := rc.QueryDefaultInt("cols", 170)
	row := rc.QueryDefaultInt("rows", 38)
	client, session, channel := core.V1.Terminal().GenerateClient(c, id, accountId, col, row)
	conn, err := global.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	global.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	core.V1.Ws().GenerateConn(conn, client, session, channel)
}
