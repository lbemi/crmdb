package sys

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mssola/useragent"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/middleware"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
)

func Login(rc *rctx.ReqCtx) {
	userForm := form.UserLoginForm{}
	rc.ShouldBind(&userForm)
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
	c := rc.Request.Request.Context()
	var registerForm form.RegisterUserForm
	rc.ShouldBind(&registerForm)
	core.V1.User().Register(c, &registerForm)
}

func Logout(rc *rctx.ReqCtx) {
	middleware.JoinBlackList(rc.Keys["token"].(*jwt.Token))
	go func() {
		defer func() {
			if r := recover(); r != nil {
				switch t := r.(type) {
				case *restfulx.OpsError:
					log.Logger.Error(t.Error())
				case error:
					log.Logger.Error(t)
				case string:
					log.Logger.Error(t)
				}
			}
		}()
		req := rc.Request.Request
		ua := useragent.New(req.UserAgent())
		bName, bVersion := ua.Browser()
		log := &logsys.LogLogin{
			Username:      rc.LoginAccount.UserName,
			Ipaddr:        rc.ClientIP(),
			LoginLocation: "",
			Browser:       bName + ":" + bVersion,
			Os:            ua.OS(),
			Platform:      ua.Platform(),
			LoginTime:     time.Now(),
			Remark:        req.UserAgent(),
		}
		log.LoginLocation = util.GetRealAddressByIP(log.Ipaddr)
		log.Status = "1"
		log.Msg = "退出登录"
		core.V1.Login().Add(log)
	}()

}

func GetUserInfoById(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.User().GetUserInfoById(c, id)
}

func GetUserList(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	condition := &sys.User{}
	pageParam := rc.GetPageQueryParam()
	condition.Status = rc.QueryParamUint8("status")
	condition.UserName = rc.Query("name")
	rc.ResData = core.V1.User().GetUserList(c, pageParam, condition)
}

func DeleteUserByUserId(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	core.V1.User().DeleteUserByUserId(c, id)
}

func UpdateUser(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var user form.UpdateUserFrom
	rc.ShouldBind(&user)
	userID := rc.PathParamUint64("id")
	core.V1.User().Update(c, userID, &user)
}

func GetUserRoles(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	uid := rc.PathParamUint64("id")
	rc.ResData = core.V1.User().GetRoleIDByUser(c, uid)
}

func SetUserRoles(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	var roles form.Roles
	rc.ShouldBind(&roles)
	uid := rc.PathParamUint64("id")
	core.V1.User().SetUserRoles(c, uid, roles.RoleIds)
}

func GetButtonsByCurrentUser(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	rc.ResData = core.V1.User().GetButtonsByUserID(c, rc.LoginAccount.ID)
}

func GetLeftMenusByCurrentUser(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	uidStr, exist := rc.Get("id")
	restfulx.ErrNotTrue(exist, restfulx.NotLogin)
	u := uidStr.(string)
	uid := util.ParseInt64(u)
	res := core.V1.User().GetLeftMenusByUserID(c, uid)
	permissions := core.V1.User().GetButtonsByUserID(c, uid)

	rc.ResData = &form.UserPermissionResp{
		Menus:      res,
		Permission: permissions,
	}
}

func UpdateUserStatus(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	userId := rc.PathParamUint64("id")
	status := rc.PathParamUint64("status")
	core.V1.User().UpdateStatus(c, userId, status)
}
