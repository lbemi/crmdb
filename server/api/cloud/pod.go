package cloud

import (
	"context"
	"io"
	"time"

	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"

	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
)

func ListPods(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	pageParam := rc.GetPageQueryParam()
	name := rc.Query("name")
	label := rc.Query("label")
	rc.ResData = core.V1.Cluster(clusterName).Pods(namespace).List(c, pageParam, name, label)
}

func GetPod(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	podName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Pods(namespace).Get(c, podName)
}

func CreatePod(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var pod *v1.Pod
	rc.ShouldBind(&pod)
	rc.ResData = core.V1.Cluster(clusterName).Pods(pod.Namespace).Create(c, pod)
}

func UpdatePod(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	clusterName := rc.QueryCloud()
	var pod *v1.Pod
	rc.ShouldBind(&pod)
	rc.ResData = core.V1.Cluster(clusterName).Pods(pod.Namespace).Update(c, pod)
}

func DeletePod(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	podName := rc.PathParam("name")
	core.V1.Cluster(clusterName).Pods(namespace).Delete(c, podName)

}
func PodExec(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	podName := rc.PathParam("name")
	containerName := rc.PathParam("container")

	conn, err := wsstore.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	wsstore.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	wsClient := wsstore.NewWsClient(conn, clusterName, "")
	err = core.V1.Cluster(clusterName).Pods(namespace).PodExec(c, namespace, podName, containerName, []string{"sh"}).StreamWithContext(c, remotecommand.StreamOptions{
		Stdout: wsClient,
		Stdin:  wsClient,
		Stderr: wsClient,
		Tty:    true,
	})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func GetPodLog(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	podName := rc.PathParam("name")
	containerName := rc.PathParam("container")

	withTimeout, cancelFunc := context.WithTimeout(c, time.Minute*10)
	defer cancelFunc()
	req := core.V1.Cluster(clusterName).Pods(namespace).GetPodLog(c, podName, containerName)
	reader, err := req.Stream(withTimeout)
	defer reader.Close()

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	conn, err := wsstore.Upgrader.Upgrade(rc.Response.ResponseWriter, rc.Request.Request, nil)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	defer conn.Close()
	wsstore.Upgrader.Subprotocols = []string{rc.Request.Request.Header.Get("Sec-WebSocket-Protocol")}
	wsClient := wsstore.NewWsClient(conn, clusterName, "log")
	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}

		err = wsClient.Conn.WriteMessage(websocket.TextMessage, buf[0:n])
		if err != nil {
			break
		}
	}
}

func GetPodEvents(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	namespace := rc.PathParam("namespace")
	clusterName := rc.QueryCloud()
	podName := rc.PathParam("name")
	rc.ResData = core.V1.Cluster(clusterName).Pods(namespace).GetPodEvent(c, podName)
}
