package v1beat1

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/websocket"
)

func Ws(rc *rctx.ReqCtx) {

	clusterName := rc.PathParam("cluster")
	typeName := rc.PathParam("type")
	websocket.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	conn, err := websocket.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
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
