package wsstore

import (
	"github.com/gorilla/websocket"
	"time"
)

type WsClient struct {
	conn     *websocket.Conn
	cluster  string
	resource string
}

func NewWsClient(conn *websocket.Conn, cluster string, resource string) *WsClient {
	return &WsClient{conn: conn, cluster: cluster, resource: resource}
}

func (w *WsClient) Ping(t time.Duration) {
	for {
		time.Sleep(t)
		err := w.conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		if err != nil {
			WsClientMap.Remove(w.conn)
			return
		}
	}
}
