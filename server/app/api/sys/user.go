package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/util"
	"strconv"
	"time"
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
		log.Logger.Error(err.Error())
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(userForm, err))
		return
	}
	//校验验证码
	//if !store.Verify(userForm.CaptchaId, userForm.Captcha, true) {
	//	response.Fail(c, 2005, "验证码错误")
	//	return
	//}

	user, err := core.Core.User().Login(c, &userForm)
	//user, err := services.Login(userForm)
	if err != nil {
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

	tokenStr, err := util.CreateToken(util.AppGuardName, user)
	if err != nil {
		log.Logger.Error(err.Error())
		response.Fail(c, response.StatusInternalServerError)
		return
	}

	if err = core.Core.Redis().Set("key", tokenStr.Token, time.Duration(time.Hour*30)); err != nil {
		log.Logger.Error(err.Error())
		response.Fail(c, response.StatusInternalServerError)
		return
	}

	response.Success(c, response.StatusOK, tokenStr)
}

func Register(c *gin.Context) {
	var registerForm form.RegisterUserForm
	if err := c.ShouldBind(&registerForm); err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrCodeParameter, util.GetErrorMsg(registerForm, err))
		return
	}

	if core.Core.User().CheckUserExist(c, registerForm.UserName) {
		response.Fail(c, response.ErrCodeUserExist)
		return
	}

	if err := core.Core.User().Register(c, &registerForm); err != nil {
		log.Logger.Error(err)
		response.FailWithMessage(c, response.ErrCodeRegisterFail, err.Error())
		return
	}

	response.Success(c, response.StatusOK, nil)

}

func Logout(c *gin.Context) {
	err := middleware.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		log.Logger.Error(err.Error())
		response.Fail(c, response.StatusInternalServerError)
		return
	}
	response.Success(c, 200, nil)
}

func GetUserInfoById(c *gin.Context) {
	id := util.GetQueryToUint64(c, "id")
	err, user := core.Core.User().GetUserInfoById(c, id)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeNotFount)
		return
	}
	response.Success(c, response.StatusOK, user)
}

func GetUserList(c *gin.Context) {
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

	user, err := core.Core.User().GetUserList(c, page, limit)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeNotFount)
		return
	}
	response.Success(c, response.StatusOK, user)
}

func DeleteUserByUserId(c *gin.Context) {

	id := util.GetQueryToUint64(c, "id")

	err := core.Core.User().DeleteUserByUserId(c, id)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func UpdateUser(c *gin.Context) {
	var (
		err  error
		user form.UpdateUserFrom
	)
	if err = c.ShouldBindJSON(&user); err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}
	userID := util.GetQueryToUint64(c, "id")

	if err = core.Core.User().Update(c, userID, &user); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func GetUserRoles(c *gin.Context) {

	uid := util.GetQueryToUint(c, "id")

	result, err := core.Core.User().GetRoleIDByUser(c, uid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, result)
}

func SetUserRoles(c *gin.Context) {
	var roles form.Roles

	err := c.ShouldBindJSON(&roles)
	if err != nil {
		response.Fail(c, response.ErrCodeParameter)
		return
	}

	uid := util.GetQueryToUint(c, "id")

	_, err = core.Core.User().GetUserInfoById(c, uid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	err = core.Core.User().SetUserRoles(c, uid, roles.RoleIds)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}
func GetButtonsByCurrentUser(c *gin.Context) {
	uidStr, exist := c.Get("userId")
	if !exist {
		response.Fail(c, response.ErrCodeNotLogin)
		return
	}
	uid := uidStr.(uint64)

	res, err := core.Core.User().GetButtonsByUserID(c, uid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}

func GetLeftMenusByCurrentUser(c *gin.Context) {
	uidStr, exist := c.Get("userId")
	if !exist {
		response.Fail(c, response.ErrCodeNotLogin)
		return
	}
	uid := uidStr.(uint64)
	res, err := core.Core.User().GetLeftMenusByUserID(c, uid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, res)
}
func UpdateUserStatus(c *gin.Context) {

	userId := util.GetQueryToUint64(c, "id")
	status := util.GetQueryToUint64(c, "status")

	if err := core.Core.User().UpdateStatus(c, userId, status); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}
