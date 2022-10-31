package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/lbemi"
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
	_, err := lbemi.CoreV1.Menu().GetMenuByMenuNameUrl(c, menu.URL, menu.Method)
	if err != gorm.ErrRecordNotFound {
		response.Fail(c, response.ErrCodeFount)
		return
	}

	if _, err := lbemi.CoreV1.Menu().Create(c, &menu); err != nil {
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

	if !lbemi.CoreV1.Menu().CheckMenusIsExist(c, menuID) {
		response.Fail(c, response.ErrCodeNotFount)
		return
	}

	if err := lbemi.CoreV1.Menu().Update(c, &menu, menuID); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func DeleteMenu(c *gin.Context) {

	menuID := util.GetQueryToUint(c, "id")

	if !lbemi.CoreV1.Menu().CheckMenusIsExist(c, menuID) {
		response.Fail(c, response.ErrCodeNotFount)
		return
	}

	if err := lbemi.CoreV1.Menu().Delete(c, menuID); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}

func GetMenu(c *gin.Context) {

	menuID := util.GetQueryToUint64(c, "id")

	res, err := lbemi.CoreV1.Menu().Get(c, menuID)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, res)
}

func ListMenus(c *gin.Context) {
	var menuType []int8
	menuTypeStr := c.DefaultQuery("menu_type", "1,2,3")
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

	res, err := lbemi.CoreV1.Menu().List(c, page, limit, menuType)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

func UpdateMenuStatus(c *gin.Context) {

	menuID := util.GetQueryToUint64(c, "id")

	status := util.GetQueryToUint64(c, "status")

	if !lbemi.CoreV1.Menu().CheckMenusIsExist(c, menuID) {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	if err := lbemi.CoreV1.Menu().UpdateStatus(c, menuID, status); err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	response.Success(c, response.StatusOK, nil)
}
