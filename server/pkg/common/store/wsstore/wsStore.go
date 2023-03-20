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
	fmt.Printf("WS-----::: %s: %s----%v ++++++++++++\n", wc.Cluster, wc.Resource, wc.Conn.RemoteAddr())
	w.data.Store(conn.RemoteAddr().String(), wc)
	go wc.Ping(time.Second * 5)
}

func (w *WsClientStore) Remove(client *websocket.Conn) {
	w.data.Delete(client.RemoteAddr().String())
}

func (w *WsClientStore) SendAll(msg interface{}) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.data.Range(func(key, value any) bool {
		c := value.(*WsClient).Conn
		fmt.Println("ws----------::::: ", key.(*WsClient).Conn)
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
		c := value.(*WsClient)
		if c.Cluster == clusterName && c.Resource == resource {
			fmt.Printf("%s--------%s 资源发生改变-----> 向《%s-%s-%s》发送数据\n", clusterName, resource, c.Cluster, c.Resource, c.Conn.RemoteAddr().String())
			err := c.Conn.WriteJSON(msg)
			if err != nil {
				w.Remove(c.Conn)
				log.Println(err)
			}
		}

		return true
	})
}
