package sys

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/common/response"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	"github.com/mssola/useragent"
	"time"
)

func Login(rc *rctx.ReqCtx) {
	userForm := form.UserLoginForm{}
	rctx.ShouldBind(rc, &userForm)
	//校验验证码
	restfulx.ErrNotTrue(store.Verify(userForm.CaptchaId, userForm.Captcha, true), restfulx.CaptchaErr)
	user := core.V1.User().Login(rc, &userForm)
	tokenStr := util.CreateToken(util.AppGuardName, user)
	core.V1.Redis().Set("key", tokenStr.Token, time.Duration(time.Hour*30))
	rc.LoginAccount = user
	rc.ResData = &form.LoginResp{
		Token: tokenStr.Token,
		User:  user,
	}
}

func Register(rc *rctx.ReqCtx) {
	var registerForm form.RegisterUserForm
	rctx.ShouldBind(rc, &registerForm)
	core.V1.User().Register(&registerForm)
}

func Logout(rc *rctx.ReqCtx) {
	middleware.JoinBlackList(rc.Keys["token"].(*jwt.Token))
	go func() {
		req := rc.Request.Request
		ua := useragent.New(req.UserAgent())
		bName, bVersion := ua.Browser()
		log := &logsys.LogLogin{
			Username:      rc.LoginAccount.UserName,
			Ipaddr:        req.RemoteAddr,
			LoginLocation: "",
			Browser:       bName + ":" + bVersion,
			Os:            ua.OS(),
			Platform:      ua.Platform(),
			LoginTime:     time.Now(),
			Remark:        req.UserAgent(),
		}
		log.Status = "1"
		log.Msg = "退出登录"
		core.V1.Login().Add(log)
	}()

}

func GetUserInfoById(c *gin.Context) {
	id := util.GetQueryToUint64(c, "id")
	err, user := core.V1.User().GetUserInfoById(c, id)
	if err != nil {
		log.Logger.Error(err)
		response.Fail(c, response.ErrCodeNotFount)
		return
	}
	response.Success(c, response.StatusOK, user)
}

func GetUserList(rc *rctx.ReqCtx) {
	pageParam := rctx.GetPageQueryParam(rc)
	rc.ResData = core.V1.User().GetUserList(pageParam)

	//if err != nil {
	//	log.Logger.Error(err)
	//	response.Fail(c, response.ErrCodeNotFount)
	//	return
	//}
	//response.Success(c, response.StatusOK, user)
}

func DeleteUserByUserId(c *gin.Context) {

	id := util.GetQueryToUint64(c, "id")

	err := core.V1.User().DeleteUserByUserId(c, id)
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

	if err = core.V1.User().Update(c, userID, &user); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}

func GetUserRoles(c *gin.Context) {

	uid := util.GetQueryToUint(c, "id")

	result, err := core.V1.User().GetRoleIDByUser(c, uid)
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

	_, err = core.V1.User().GetUserInfoById(c, uid)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	err = core.V1.User().SetUserRoles(c, uid, roles.RoleIds)
	if err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}
	response.Success(c, response.StatusOK, nil)
}
func GetButtonsByCurrentUser(c *gin.Context) {
	uidStr, exist := c.Get("id")
	if !exist {
		response.Fail(c, response.ErrCodeNotLogin)
		return
	}
	u := uidStr.(string)
	uid := util.ParseInt64(u)

	res := core.V1.User().GetButtonsByUserID(uid)

	response.Success(c, response.StatusOK, res)
}

func GetLeftMenusByCurrentUser(rc *rctx.ReqCtx) {
	uidStr, exist := rc.Get("id")
	restfulx.ErrNotTrue(exist, restfulx.NotLogin)
	u := uidStr.(string)
	uid := util.ParseInt64(u)
	res := core.V1.User().GetLeftMenusByUserID(uid)
	permissions := core.V1.User().GetButtonsByUserID(uid)

	rc.ResData = &form.UserPermissionResp{
		Menus:      res,
		Permission: permissions,
	}
}
func UpdateUserStatus(c *gin.Context) {

	userId := util.GetQueryToUint64(c, "id")
	status := util.GetQueryToUint64(c, "status")

	if err := core.V1.User().UpdateStatus(c, userId, status); err != nil {
		response.Fail(c, response.ErrOperateFailed)
		return
	}

	response.Success(c, response.StatusOK, nil)
}
