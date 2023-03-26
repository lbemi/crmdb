package cloud

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/handler/types"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/remotecommand"
	"net/http"
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

	for {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		w, err := c.Writer.Write([]byte(string(buf[0:n])))
		if w == 0 || err != nil {
			break
		}
		c.Writer.(http.Flusher).Flush()
	}
}
