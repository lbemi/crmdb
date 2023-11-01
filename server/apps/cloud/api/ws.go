package api

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func Ws(rc *rctx.ReqCtx) {
	clusterName := rc.PathParam("cluster")
	typeName := rc.PathParam("type")
	global.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	conn, err := global.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	if err != nil {
		log.Logger.Error(err)
	} else {
		cache.WsClientMap.Store(clusterName, typeName, conn)
	}
}

func WsSendAll(rc *rctx.ReqCtx) {
	msg := rc.Query("msg")
	cache.WsClientMap.SendAll(msg)
}
