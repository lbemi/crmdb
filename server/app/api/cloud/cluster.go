package cloud

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/util"
	"io"
)

func CreateCluster(c *gin.Context) {

	//file, header, err := c.Request.FormFile("file")
	fileHeader, err := c.FormFile("file")
	util.GinError(c, err, response.ErrCodeParameter)
	file, err := fileHeader.Open()
	util.GinError(c, err, response.ErrCodeParameter)
	bytes, err := io.ReadAll(file)
	util.GinError(c, err, response.ErrCodeParameter)

	var req form.ClusterReq
	err = c.ShouldBindUri(&req)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	req.Name = c.PostForm("name")

	req.KubeConfig = string(bytes)
	err = core.V1.Cluster(req.Name).Create(c, &req)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func ListCluster(c *gin.Context) {
	list, err := core.V1.Cluster("").List(c)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, list)
}

func DeleteCluster(c *gin.Context) {
	id := util.GetQueryToUint64(c, "id")

	err := core.V1.Cluster("").Delete(c, id)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}
