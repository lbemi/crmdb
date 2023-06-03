package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
)

func Ws(c *gin.Context) {
	clusterName := c.Param("cluster")
	typeName := c.Param("type")
	conn, err := wsstore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Logger.Error(err)
	} else {
		wsstore.WsClientMap.Store(clusterName, typeName, conn)
	}
}

func WsSendAll(c *gin.Context) {
	msg := c.Query("msg")
	wsstore.WsClientMap.SendAll(msg)
}