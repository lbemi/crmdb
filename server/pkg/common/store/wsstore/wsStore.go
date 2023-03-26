package wsstore

import (
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"sync"
	"time"
)

var WsClientMap *WsClientStore

func init() {
	WsClientMap = &WsClientStore{}
}

var wsLock sync.Mutex

type WsClientStore struct {
	data sync.Map
	lock sync.Mutex
}

func (w *WsClientStore) Store(cluster, resource string, conn *websocket.Conn) {
	wc := NewWsClient(conn, cluster, resource)
	w.data.Store(conn.RemoteAddr().String(), wc)
	go wc.Ping(time.Second * 3)
}

func (w *WsClientStore) Remove(client *websocket.Conn) {
	w.data.Delete(client.RemoteAddr().String())
}

func (w *WsClientStore) SendAll(msg interface{}) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.data.Range(func(key, value any) bool {
		c := value.(*WsClient).Conn
		err := c.WriteJSON(msg)
		if err != nil {
			w.Remove(c)
			log.Println(err)
		}
		return true
	})
}

func (w *WsClientStore) SendClusterResource(clusterName, resource string, msg interface{}) {
	closeCh := make(chan struct{})
	defer close(closeCh)

	w.data.Range(func(key, value any) bool {
		c := value.(*WsClient)
		resourceName := strings.Split(c.Resource, ",")
		for _, name := range resourceName {
			if c.Cluster == clusterName && name == resource {
				wsLock.Lock()
				defer wsLock.Unlock()
				err := c.Conn.WriteJSON(msg)

				if err != nil {
					log.Println(err)
					w.Remove(c.Conn)
				}

			}
		}

		return true
	})

}
