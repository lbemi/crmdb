package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
	"net/http"
)

func Login(c *gin.Context) {
	userForm := form.UserLoginForm{}
	if err := c.ShouldBind(&userForm); err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, 201, util.GetErrorMsg(userForm, err))
		return
	}
	user, err := services.Login(userForm)
	if err != nil {
		response.Fail(c, 2001, err.Error())
		return
	}
	tokenStr, err := services.CreateToken(services.AppGuardName, user)
	if err != nil {
		response.Fail(c, 2002, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "登录成功", tokenStr)
}
func Logout(c *gin.Context) {
	c.String(http.StatusOK, "Logout")
}

func GetUserInfoById(c *gin.Context) {
	err, user := services.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.Fail(c, 2003, err.Error())
		return
	}
	response.Success(c, 200, "", user)
}
