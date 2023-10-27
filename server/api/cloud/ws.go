package cloud

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func Ws(rc *rctx.ReqCtx) {
	clusterName := rc.PathParam("cluster")
	typeName := rc.PathParam("type")
	wsstore.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	conn, err := wsstore.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	if err != nil {
		log.Logger.Error(err)
	} else {
		wsstore.WsClientMap.Store(clusterName, typeName, conn)
	}
}

func WsSendAll(rc *rctx.ReqCtx) {
	msg := rc.Query("msg")
	wsstore.WsClientMap.SendAll(msg)
}
