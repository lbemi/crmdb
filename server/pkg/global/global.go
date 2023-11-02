package global

import (
	"github.com/gorilla/websocket"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"net/http"
)

var (
	Logger   log.LoggerInterface
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}
)
