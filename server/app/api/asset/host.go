package asset

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/lbemi"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"net/http"
	"strconv"
)

// @Summary      Add host
// @Description  Add host
// @Tags         host
// @Accept       json
// @Produce      json
// @Param        data body asset.HostReq true "host info"
// @Success      200  {object}  httputils.HttpOK
// @Failure      400  {object}  httputils.HttpError
// @Router       /hosts [post]

func AddHost(c *gin.Context) {
	var machine asset.HostReq
	if err := c.ShouldBindJSON(&machine); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if err := lbemi.CoreV1.Host().Create(c, &machine); err != nil {
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

	res, err := lbemi.CoreV1.Host().List(c, page, limit)
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
	res, err := lbemi.CoreV1.Host().GetByHostId(c, id)
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

	if !lbemi.CoreV1.Host().CheckHostExist(c, id) {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	if err := lbemi.CoreV1.Host().Update(c, id, &machine); err != nil {
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

	if err := lbemi.CoreV1.Host().Delete(c, id); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsShell(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	if !lbemi.CoreV1.Host().CheckHostExist(c, id) {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	cols := c.DefaultQuery("cols", "150")
	rows := c.DefaultQuery("rows", "35")
	col, _ := strconv.Atoi(cols)
	row, _ := strconv.Atoi(rows)

	client, session, channel, err := lbemi.CoreV1.Terminal().GenerateClient(c, id, col, row)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	err = lbemi.CoreV1.Ws().GenerateConn(conn, client, session, channel)

	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
}
