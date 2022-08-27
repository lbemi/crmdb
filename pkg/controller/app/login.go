package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
	"net/http"
)

func Login(c *gin.Context) {
	userForm := form.UserLoginForm{}
	if err := c.ShouldBindJSON(&userForm); err != nil {
		response.Fail(c, 201, util.GetErrorMsg(userForm, err))
		return
	}
	_, err := services.Login(userForm)
	if err != nil {
		response.Fail(c, 201, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "登录成功", "")
}
func Logout(c *gin.Context) {
	c.String(http.StatusOK, "Logout")
}
