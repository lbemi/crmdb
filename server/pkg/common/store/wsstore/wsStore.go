package wsstore

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var WsClientMap *WsClientStore

func init() {
	WsClientMap = &WsClientStore{}
}

type WsClientStore struct {
	data sync.Map
}

func (w *WsClientStore) Store(conn *websocket.Conn) {
	wc := NewWsClient(conn)
	w.data.Store(conn.RemoteAddr().String(), wc)
	go wc.Ping(time.Second * 5)
}

func (w *WsClientStore) Remove(client *websocket.Conn) {
	w.data.Delete(client.RemoteAddr().String())
}

func (w *WsClientStore) SendAll(msg string) {
	w.data.Range(func(key, value any) bool {
		c := value.(*WsClient).conn
		err := c.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			w.Remove(c)
			log.Println(err)
		}
		return true
	})
}
