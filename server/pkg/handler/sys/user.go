package sys

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
	"github.com/mssola/useragent"
	"time"
)

type UserGetter interface {
	User() IUSer
}

type IUSer interface {
	Login(rc *rctx.ReqCtx, params *form.UserLoginForm) (user *sys.User)
	Register(params *form.RegisterUserForm)
	Update(userID uint64, params *form.UpdateUserFrom)
	GetUserInfoById(id uint64) (user *sys.User)
	GetUserList(param *model.PageParam, condition *sys.User) *form.PageUser
	DeleteUserByUserId(id uint64)
	CheckUserExist(userName string) bool
	GetByName(name string) *sys.User
	GetRoleIDByUser(userID uint64) *[]sys.Role
	SetUserRoles(userID uint64, roleIDs []uint64)
	GetButtonsByUserID(userID uint64) *[]string
	GetLeftMenusByUserID(userID uint64) *[]sys.Menu
	UpdateStatus(userID, status uint64)
}

type user struct {
	factory services.FactoryImp
}

func NewUser(f services.FactoryImp) IUSer {
	return &user{
		factory: f,
	}
}

func (u *user) Login(rc *rctx.ReqCtx, params *form.UserLoginForm) (user *sys.User) {
	user, err := u.factory.User().Login(params)
	restfulx.ErrNotNilDebug(err, restfulx.PasswdWrong)
	pass := util.BcryptMakeCheck([]byte(params.Password), user.Password)
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
			Username:      params.UserName,
			Ipaddr:        req.RemoteAddr,
			LoginLocation: "",
			Browser:       bName + ":" + bVersion,
			Os:            ua.OS(),
			Platform:      ua.Platform(),
			LoginTime:     time.Now(),
			Remark:        req.UserAgent(),
		}
		if pass && err == nil {
			log.Status = "1"
			log.Msg = "登录成功"
		} else {
			log.Status = "-1"
			log.Msg = "登录失败"
		}
		u.factory.Log().Add(log)
	}()
	restfulx.ErrNotTrue(user.Status == 1, restfulx.UserDeny)
	restfulx.ErrNotTrue(pass, restfulx.PasswdWrong)
	return user
}

func (u *user) Register(params *form.RegisterUserForm) {
	userInfo := &sys.User{
		UserName: params.UserName,
		Password: util.BcryptMake([]byte(params.Password)),
		//Mobile:      params.Mobile,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}

	restfulx.ErrNotTrue(!u.factory.User().CheckUserExist(userInfo.UserName), restfulx.UserExist)

	restfulx.ErrNotNilDebug(u.factory.User().Register(userInfo), restfulx.OperatorErr)
}

func (u *user) Update(userID uint64, params *form.UpdateUserFrom) {
	userInfo := &sys.User{
		UserName:    params.UserName,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}
	restfulx.ErrNotNilDebug(u.factory.User().Update(userID, userInfo), restfulx.OperatorErr)
}

func (u *user) GetUserInfoById(id uint64) *sys.User {
	res, err := u.factory.User().GetUserInfoById(id)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (u *user) GetUserList(pageParam *model.PageParam, condition *sys.User) *form.PageUser {
	res, err := u.factory.User().GetUserList(pageParam, condition)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (u *user) DeleteUserByUserId(id uint64) {
	err := u.factory.User().DeleteUserByUserId(id)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	u.factory.Authentication().DeleteUser(id)
}

func (u *user) CheckUserExist(userName string) bool {
	return u.factory.User().CheckUserExist(userName)
}

func (u *user) GetByName(name string) *sys.User {
	res, err := u.factory.User().GetByName(name)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

// GetRoleIDByUser 查询用户角色
func (u *user) GetRoleIDByUser(userID uint64) *[]sys.Role {
	res, err := u.factory.User().GetRoleIdbyUser(userID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

// SetUserRoles 分配用户角色
func (u *user) SetUserRoles(userID uint64, roleIDS []uint64) {

	// 配置role_users表
	tx, err := u.factory.User().SetUserRoles(userID, roleIDS)
	if err != nil {
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	// 添加规则到rules表
	err = u.factory.Authentication().AddRoleForUser(userID, roleIDS)
	if err != nil {
		// 报错则回退数据
		tx.Rollback()
		for _, roleId := range roleIDS {
			u.factory.Authentication().DeleteRoleWithUser(userID, roleId)
		}
	}

}

// GetButtonsByUserID 获取菜单按钮
func (u *user) GetButtonsByUserID(userID uint64) *[]string {
	menus, err := u.factory.User().GetButtonsByUserID(userID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	var res []string
	for _, v := range *menus {
		if v.Code != "" {
			res = append(res, v.Code)
		}
	}

	return &res
}

// GetLeftMenusByUserID 根据用户ID获取左侧菜单
func (u *user) GetLeftMenusByUserID(userID uint64) *[]sys.Menu {
	res, err := u.factory.User().GetLeftMenusByUserID(userID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (u *user) UpdateStatus(userId, status uint64) {
	restfulx.ErrNotNilDebug(u.factory.User().UpdateStatus(userId, status), restfulx.OperatorErr)
}
