package cloud

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"github.com/lbemi/lbemi/pkg/util"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"
	"strconv"
	"time"
)

func ListPods(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}
	podList, err := core.V1.Cluster(clusterName).Pods(namespace).List(c)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(podList)

	if pageQuery.Total <= limit {
		pageQuery.Data = podList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = podList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = podList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

func GetPod(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")

	podName := c.Param("podName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	pod, err := core.V1.Cluster(clusterName).Pods(namespace).Get(c, podName)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, pod)
}

func CreatePod(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var pod *v1.Pod

	err := c.ShouldBindJSON(&pod)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newPod, err := core.V1.Cluster(clusterName).Pods(pod.Namespace).Create(c, pod)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newPod)
}

func UpdatePod(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var pod *v1.Pod

	err := c.ShouldBindJSON(&pod)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	newPod, err := core.V1.Cluster(clusterName).Pods(pod.Namespace).Update(c, pod)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, newPod)
}

func DeletePod(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	podName := c.Param("podName")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	err := core.V1.Cluster(clusterName).Pods(namespace).Delete(c, podName)
	if err != nil {
		response.FailWithMessage(c, response.ErrOperateFailed, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)
}
func PodExec(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	podName := c.Param("podName")
	containerName := c.Param("container")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	conn, err := wsstore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	wsClient := wsstore.NewWsClient(conn, clusterName, "")
	err = core.V1.Cluster(clusterName).Pods(namespace).PodExec(c, namespace, podName, containerName, []string{"sh"}).StreamWithContext(c, remotecommand.StreamOptions{
		Stdout: wsClient,
		Stdin:  wsClient,
		Stderr: wsClient,
		Tty:    true,
	})
	if err != nil {
		log.Logger.Error("-------", err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
}

func GetPodLog(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	podName := c.Param("podName")
	containerName := c.Param("container")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	withTimeout, cancelFunc := context.WithTimeout(c, time.Minute*10)
	defer cancelFunc()
	req := core.V1.Cluster(clusterName).Pods(namespace).GetPodLog(c, podName, containerName)
	reader, err := req.Stream(withTimeout)
	if err != nil {
		response.Fail(c, response.StatusInternalServerError)
		return
	}
	defer reader.Close()

	conn, err := wsstore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	defer conn.Close()

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

func GetPodEvents(c *gin.Context) {
	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	namespace := c.Param("namespace")
	podName := c.Param("name")
	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	events, err := core.V1.Cluster(clusterName).Pods(namespace).GetPodEvent(c, podName)
	util.GinError(c, err, response.ErrCodeParameter)

	response.Success(c, response.StatusOK, events)
}

func SearchPods(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	limitStr := c.DefaultQuery("limit", "0")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	clusterName := c.Query("cloud")
	if clusterName == "" {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	namespace := c.Param("namespace")

	if !core.V1.Cluster(clusterName).CheckHealth(c) {
		response.Fail(c, response.ClusterNoHealth)
		return
	}

	key := c.DefaultQuery("key", "")
	searchTypeStr := c.DefaultQuery("type", "0")
	searchType, err := strconv.Atoi(searchTypeStr)

	podList, err := core.V1.Cluster(clusterName).Pods(namespace).Search(c, key, searchType)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	// 处理分页
	var pageQuery types.PageQuery
	pageQuery.Total = len(podList)

	if pageQuery.Total <= limit {
		pageQuery.Data = podList
	} else if page*limit >= pageQuery.Total {
		pageQuery.Data = podList[(page-1)*limit : pageQuery.Total]
	} else {
		pageQuery.Data = podList[(page-1)*limit : page*limit]
	}

	response.Success(c, response.StatusOK, pageQuery)
}

/*
// WebSocket 读取 Pod 实时日志并推送给前端
func tailLog(podName, containerName string, ws *websocket.Conn) {
    // 创建 Kubernetes REST API 客户端
    clientset, err := kubernetes.NewForConfig(kubeconfig)
    if err != nil {
        log.Printf("Failed to create Kubernetes clientset: %v", err)
        return
    }

    // 获取 Pod 对象
    pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
    if err != nil {
        log.Printf("Failed to get Pod %s: %v", podName, err)
        return
    }

    // 通过 Pod 对象获取日志
    req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
        Container: containerName,
        Follow:    true,
    })
    logStream, err := req.Stream(context.Background())
    if err != nil {
        log.Printf("Failed to open log stream for Pod %s container %s: %v", podName, containerName, err)
        return
    }
    defer logStream.Close()

    // 读取日志并发送到 WebSocket
    buf := make([]byte, 1024)
    for {
        // 读取日志
        n, err := logStream.Read(buf)
        if err != nil {
            log.Printf("Failed to read log for Pod %s container %s: %v", podName, containerName, err)
            return
        }

        // 发送到 WebSocket
        err = ws.WriteMessage(websocket.TextMessage, buf[:n])
        if err != nil {
            log.Printf("Failed to write log to WebSocket for Pod %s container %s: %v", podName, containerName, err)
            return
        }
    }
}

// WebSocket 处理函数
func handleWebSocket(c *gin.Context) {
    // 升级为 WebSocket 连接
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Failed to upgrade to WebSocket: %v", err)
        return
    }
    defer ws.Close()

    // 从 URL 参数中获取 Pod 名称和容器名称
    podName := c.Query("pod")
    containerName := c.Query("container")

    // 读取 Pod 实时日志并推送给前端
    tailLog(podName, containerName, ws)
}

*/
