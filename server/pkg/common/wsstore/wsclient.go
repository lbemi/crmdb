package wsstore

//
//import (
//	"github.com/lbemi/lbemi/pkg/cache"
//	"time"
//
//	"github.com/gorilla/websocket"
//)
//
//type WsClient struct {
//	Conn     *websocket.Conn
//	Cluster  string
//	Resource string
//}
//
//func NewWsClient(conn *websocket.Conn, cluster string, resource string) *WsClient {
//	return &WsClient{Conn: conn, Cluster: cluster, Resource: resource}
//}
//
//func (w *WsClient) Ping(t time.Duration) {
//	for {
//		time.Sleep(t)
//		cache.WsLock.Lock()
//		err := w.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
//		cache.WsLock.Unlock()
//		if err != nil {
//			cache.WsClientMap.Remove(w.Conn)
//			return
//		}
//	}
//}
//
//func (w *WsClient) Write(p []byte) (n int, err error) {
//	err = w.Conn.WriteMessage(websocket.TextMessage, p)
//	if err != nil {
//		return 0, err
//	}
//
//	return len(p), nil
//}
//func (w *WsClient) Read(p []byte) (n int, err error) {
//	_, bytes, err := w.Conn.ReadMessage()
//	if err != nil {
//		return 0, err
//	}
//	return copy(p, string(bytes)), nil
//}
