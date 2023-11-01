package cache

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var WsClientMap *WsClientStore

func init() {
	WsClientMap = &WsClientStore{}
}

var WsLock sync.Mutex

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
	fmt.Println("发送消息啦。。。。。。。")
	w.data.Range(func(key, value any) bool {
		c := value.(*WsClient)
		resourceName := strings.Split(c.Resource, ",")
		for _, name := range resourceName {
			if c.Cluster == clusterName && name == resource {
				WsLock.Lock()
				defer WsLock.Unlock()
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
		WsLock.Lock()
		err := w.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		WsLock.Unlock()
		if err != nil {
			WsClientMap.Remove(w.Conn)
			return
		}
	}
}

func (w *WsClient) Write(p []byte) (n int, err error) {
	err = w.Conn.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}
func (w *WsClient) Read(p []byte) (n int, err error) {
	_, bytes, err := w.Conn.ReadMessage()
	if err != nil {
		return 0, err
	}
	return copy(p, string(bytes)), nil
}
