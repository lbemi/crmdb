package api

import (
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func WebSocket(rc *rctx.ReqCtx) {
	clusterName := rc.PathParam("cluster")
	typeName := rc.PathParam("type")
	global.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	conn, err := global.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	if err != nil {
		global.Logger.Error(err)
	} else {
		cache.WebsocketStore.Store(clusterName, typeName, conn)
	}
}

func WsSendAll(rc *rctx.ReqCtx) {
	msg := rc.Query("msg")
	cache.WebsocketStore.SendAll(msg)
}
