package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
)

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 登录
// @Accept json
// @Produce  json
// @Param data body form.UserLoginForm true "Form表单"
// @Success 200 {object} response.Response{}  "请求成功"
// @Failure 2005 {object} response.Response{data=util.TokenOutPut} "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /login [post]
func Login(c *gin.Context) {
	userForm := form.UserLoginForm{}
	if err := c.ShouldBind(&userForm); err != nil {
		global.App.Log.Error(err.Error())
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(userForm, err))
		return
	}
	//校验验证码
	//if !store.Verify(userForm.CaptchaId, userForm.Captcha, true) {
	//	response.Fail(c, 2005, "验证码错误")
	//	return
	//}

	user, err := services.Login(userForm)
	if user.ID == 0 {
		response.Fail(c, response.ErrCodeUserNotExist)
		return
	}
	if user.Status != 1 {
		response.Fail(c, response.ErrCodeUserForbidden)
		return
	}
	if ok := util.BcryptMakeCheck([]byte(userForm.Password), user.Password); !ok {
		response.Fail(c, response.ErrCodeUserOrPasswdWrong)
		return
	}
	if err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.ErrCodeUserOrPasswdWrong)
		return
	}
	tokenStr, err := util.CreateToken(util.AppGuardName, user)
	if err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.StatusInternalServerError)
		return
	}
	global.App.Redis.Set("key", tokenStr.Token, time.Duration(time.Hour*30))
	response.Success(c, response.LoginSuccess, tokenStr)
}

func Logout(c *gin.Context) {
	err := util.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.StatusInternalServerError)
		return
	}
	response.Success(c, 200, nil)
}

func GetUserInfoById(c *gin.Context) {
	err, user := services.GetUserInfoById(c.Keys["id"].(string))
	if err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.ErrCodeNotFount)
		return
	}
	response.Success(c, response.StatusOK, user)
}

func GetUserList(c *gin.Context) {
	err, user := services.GetUserList()
	if err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.ErrCodeNotFount)
		return
	}
	response.Success(c, response.StatusOK, user)
}

func DeleteUserByUserId(c *gin.Context) {

	id := util.GetQueryToUint64(c, "id")

	err := services.DeleteUserByUserId(id)
	if err != nil {
		global.App.Log.Error(err.Error())
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}