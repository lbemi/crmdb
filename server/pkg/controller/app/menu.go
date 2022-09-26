package app

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/sys"
)

type Menu struct{}

func GetMenuList(c *gin.Context) {
	var menus []sys.Menu
	err := global.App.DB.Model(&sys.Menu{}).Order("parent_id ASC").Order("sequence ASC").Find(&menus).Error
	if err != nil {
		global.App.Log.Error(err.Error())
		return
	}
	global.App.Log.Info("menus", zap.Any("menus", menus))
	response.Success(c, response.StatusOK, menus)
}

// 获取菜单有权限的操作列表
//func (Menu) MenuButtonList(c *gin.Context) {
//	// 用户ID
//	uid, isExit := c.Get("id")
//	if !isExit {
//		response.Fail(c, response.InvalidToken)
//		return
//	}
//	//userID :=
//	menuCode := common.GetQueryToStr(c, "menucode")
//	if userID == 0 || menuCode == "" {
//		common.ResFail(c, "err")
//		return
//	}
//	btnList := []string{}
//	if userID == common.SUPER_ADMIN_ID {
//		//管理员
//		btnList = append(btnList, "add")
//		btnList = append(btnList, "del")
//		btnList = append(btnList, "view")
//		btnList = append(btnList, "update")
//		btnList = append(btnList, "setrolemenu")
//		btnList = append(btnList, "setadminrole")
//	} else {
//		menu := sys.Menu{}
//		err := menu.GetMenuButton(userID, menuCode, &btnList)
//		if err != nil {
//			common.ResErrSrv(c, err)
//			return
//		}
//	}
//	common.ResSuccess(c, &btnList)
//}
