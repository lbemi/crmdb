package wsstore

import (
	"fmt"
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
	lock sync.Mutex
}

func (w *WsClientStore) Store(cluster, resource string, conn *websocket.Conn) {
	wc := NewWsClient(conn, cluster, resource)
	w.data.Store(wc, conn.RemoteAddr().String())
	go wc.Ping(time.Second * 5)
}

func (w *WsClientStore) Remove(client *websocket.Conn) {
	w.data.Delete(client.RemoteAddr().String())
}

func (w *WsClientStore) SendAll(msg interface{}) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.data.Range(func(key, value any) bool {
		c := key.(*WsClient).conn
		fmt.Println("ws----------::::: ", key.(*WsClient).conn)
		err := c.WriteJSON(msg)
		if err != nil {
			w.Remove(c)
			log.Println(err)
		}
		return true
	})
}

func (w *WsClientStore) SendClusterResource(clusterName, resource string, msg interface{}) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.data.Range(func(key, value any) bool {
		c := key.(*WsClient).conn
		if value.(*WsClient).cluster == clusterName && value.(*WsClient).resource == resource {
			fmt.Println("ws----------::::: ", key.(*WsClient).conn)
			err := c.WriteJSON(msg)
			if err != nil {
				w.Remove(c)
				log.Println(err)
			}
		}

		return true
	})
}
