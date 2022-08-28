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

// @Summary 用户登录
// @Description 用户登录
// @Accept json
// @Produce  json
// @Param data body form.UserLoginForm true "请示参数data"
// @Success 200 {object} response.Response
//"请求成功"
// @Failure 2005 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /login [post]
func Login(c *gin.Context) {
	userForm := form.UserLoginForm{}
	if err := c.ShouldBind(&userForm); err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, 201, util.GetErrorMsg(userForm, err))
		return
	}
	//校验验证码
	if !store.Verify(userForm.CaptchaId, userForm.Captcha, true) {
		response.Fail(c, 2005, "验证码错误")
		return
	}

	user, err := services.Login(userForm)
	if err != nil {
		response.Fail(c, 2001, err.Error())
		return
	}
	tokenStr, err := util.CreateToken(util.AppGuardName, user)
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
	err, user := services.GetUserInfoById(c.Keys["id"].(string))
	if err != nil {
		response.Fail(c, 2003, err.Error())
		return
	}
	response.Success(c, 200, "", user)
}
func GetUserInfos(c *gin.Context) {
	err, user := services.GetUserInfos(c.Keys["id"].(string))
	if err != nil {
		response.Fail(c, 2003, err.Error())
		return
	}
	response.Success(c, 200, "", user)
}