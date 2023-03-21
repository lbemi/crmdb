package wsstore

import (
	"github.com/gorilla/websocket"
	"time"
)

type WsClient struct {
	Conn     *websocket.Conn
	Cluster  string
	Resource string
}

func NewWsClient(conn *websocket.Conn, cluster string, resource string) *WsClient {
	return &WsClient{Conn: conn, Cluster: cluster, Resource: resource}
}

func (w *WsClient) Ping(t time.Duration) {
	for {
		time.Sleep(t)
		wsLock.Lock()
		err := w.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		wsLock.Unlock()
		if err != nil {
add			WsClientMap.Remove(w.Conn)
			return
		}

	}
}
