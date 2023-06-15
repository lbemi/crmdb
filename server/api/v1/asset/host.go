package asset

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/common/store/wsstore"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/asset"
)

func AddHost(c *gin.Context) {
	var machine asset.HostReq
	if err := c.ShouldBindJSON(&machine); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if err := core.V1.Host().Create(c, &machine); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)

}

// @Summary      List host
// @Description  List host
// @Tags         host
// @Accept       json
// @Produce      json
// @Success      200  {object}  httputils.Response{result=asset.Host}
// @Failure      400  {object}  httputils.HttpError
// @Router       /hosts [get]

func ListHosts(c *gin.Context) {
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

	res, err := core.V1.Host().List(c, page, limit)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

// @Summary      Get host by id
// @Description  Get host by id
// @Tags         host
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "host ID"
// @Success      200  {object}  httputils.Response{result=asset.Host}
// @Failure      400  {object}  httputils.HttpError
// @Router       /hosts/:id [get]

func GetHostById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	res, err := core.V1.Host().GetByHostId(c, id)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

// @Summary      Get host by id
// @Description  Get host by id
// @Tags         host
// @Accept       json
// @Produce      json
// @Param        data body asset.HostReq true "host info"
// @Param        id   path      int  true  "host ID"
// @Success      200  {object}  httputils.Response{result=asset.Host}
// @Failure      400  {object}  httputils.HttpError
// @Router       /hosts/:id [put]

func UpdateHost(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	var machine asset.HostReq
	if err := c.ShouldBindJSON(&machine); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if !core.V1.Host().CheckHostExist(c, id) {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	if err := core.V1.Host().Update(c, id, &machine); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func DeleteHost(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if err := core.V1.Host().Delete(c, id); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func WsShell(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if !core.V1.Host().CheckHostExist(c, id) {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	cols := c.DefaultQuery("cols", "170")
	rows := c.DefaultQuery("rows", "38")
	col, _ := strconv.Atoi(cols)
	row, _ := strconv.Atoi(rows)

	client, session, channel, err := core.V1.Terminal().GenerateClient(c, id, col, row)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	conn, err := wsstore.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	err = core.V1.Ws().GenerateConn(conn, client, session, channel)

	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
}
