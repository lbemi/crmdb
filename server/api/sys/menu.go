package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func AddMenu(c *gin.Context) {

	var menu form.MenusReq
	if err := c.ShouldBindJSON(&menu); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	// 判断权限是否已存在
	_, err := core.V1.Menu().GetMenuByMenuNameUrl(c, menu.Path, menu.Method)
	if err != gorm.ErrRecordNotFound {
		response.Fail(c, response.ErrCodeFount)
		return
	}

	if _, err := core.V1.Menu().Create(c, &menu); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func UpdateMenu(c *gin.Context) {

	var menu form.UpdateMenusReq

	if err := c.ShouldBindJSON(&menu); err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	menuID := util.GetQueryToUint64(c, "id")

	if !core.V1.Menu().CheckMenusIsExist(c, menuID) {
		response.Fail(c, response.ErrCodeNotFount)
		return
	}

	if err := core.V1.Menu().Update(c, &menu, menuID); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func DeleteMenu(c *gin.Context) {

	menuID := util.GetQueryToUint(c, "id")

	if !core.V1.Menu().CheckMenusIsExist(c, menuID) {
		response.Fail(c, response.ErrCodeNotFount)
		return
	}

	if err := core.V1.Menu().Delete(c, menuID); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func GetMenu(c *gin.Context) {

	menuID := util.GetQueryToUint64(c, "id")

	res, err := core.V1.Menu().Get(c, menuID)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, res)
}

func ListMenus(c *gin.Context) {
	var menuType []int8
	isTree := true

	menuTypeStr := c.DefaultQuery("menuType", "1,2,3")
	tree := c.DefaultQuery("isTree", "true")

	if tree == "false" {
		isTree = false
	}

	menuTypeSlice := strings.Split(menuTypeStr, ",")
	for _, t := range menuTypeSlice {
		res, err := strconv.Atoi(t)
		if err != nil {
			response.Fail(c, response.ErrCodeParameter)
			return
		}
		menuType = append(menuType, int8(res))
	}

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

	res, err := core.V1.Menu().List(c, page, limit, menuType, isTree)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

func UpdateMenuStatus(c *gin.Context) {

	menuID := util.GetQueryToUint64(c, "id")

	status := util.GetQueryToUint64(c, "status")

	if !core.V1.Menu().CheckMenusIsExist(c, menuID) {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	if err := core.V1.Menu().UpdateStatus(c, menuID, status); err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	response.Success(c, response.StatusOK, nil)
}
