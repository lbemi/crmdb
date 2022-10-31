package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/lbemi"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/util"
	"strconv"
)

func AddRole(c *gin.Context) {
	var role form.RoleReq
	if err := c.ShouldBind(&role); err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(role, err))
		return
	}

	exist := lbemi.CoreV1.Role().CheckRoleIsExist(c, role.Name)
	if exist {
		response.Fail(c, response.ErrCodeFount)
		return
	}

	if _, err := lbemi.CoreV1.Role().Create(c, &role); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func UpdateRole(c *gin.Context) {

	var role form.UpdateRoleReq
	if err := c.ShouldBindJSON(&role); err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(role, err))
		return
	}

	roleID := util.GetQueryToUint64(c, "id")

	_, err := lbemi.CoreV1.Role().Get(c, roleID)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	if err = lbemi.CoreV1.Role().Update(c, &role, roleID); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func DeleteRole(c *gin.Context) {

	rid := util.GetQueryToUint64(c, "id")

	_, err := lbemi.CoreV1.Role().Get(c, rid)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	if err = lbemi.CoreV1.Role().Delete(c, rid); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func GetRole(c *gin.Context) {

	rid := util.GetQueryToUint64(c, "id")

	res, err := lbemi.CoreV1.Role().Get(c, rid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

func ListRoles(c *gin.Context) {

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

	res, err := lbemi.CoreV1.Role().List(c, page, limit)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

func GetMenusByRole(c *gin.Context) {

	rid := util.GetQueryToUint(c, "id")

	_, err := lbemi.CoreV1.Role().Get(c, rid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	res, err := lbemi.CoreV1.Role().GetMenusByRoleID(c, rid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

func SetRoleMenus(c *gin.Context) {

	rid := util.GetQueryToUint(c, "id")

	_, err := lbemi.CoreV1.Role().Get(c, rid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	var menuIDs form.Menus
	if err = c.ShouldBindJSON(&menuIDs); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	if err = lbemi.CoreV1.Role().SetRole(c, rid, menuIDs.MenuIDS); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func UpdateRoleStatus(c *gin.Context) {

	status := util.GetQueryToUint64(c, "status")

	roleID := util.GetQueryToUint64(c, "id")

	_, err := lbemi.CoreV1.Role().Get(c, roleID)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	if err = lbemi.CoreV1.Role().UpdateStatus(c, roleID, status); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}
