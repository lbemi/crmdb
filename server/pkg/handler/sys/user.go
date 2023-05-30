package sys

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/services"
	"github.com/lbemi/lbemi/pkg/util"
)

type UserGetter interface {
	User() IUSer
}

type IUSer interface {
	Login(params *form.UserLoginForm) (user *sys.User)
	Register(params *form.RegisterUserForm)
	Update(c context.Context, userID uint64, params *form.UpdateUserFrom) (err error)
	GetUserInfoById(c context.Context, id uint64) (user *sys.User, err error)
	GetUserList(param *model.PageParam) *form.PageUser
	DeleteUserByUserId(c context.Context, id uint64) error
	CheckUserExist(c context.Context, userName string) bool
	GetByName(c context.Context, name string) (*sys.User, error)
	GetRoleIDByUser(c context.Context, userID uint64) (*[]sys.Role, error)
	SetUserRoles(c context.Context, userID uint64, roleIDs []uint64) error
	GetButtonsByUserID(userID uint64) *[]string
	GetLeftMenusByUserID(userID uint64) *[]sys.Menu
	UpdateStatus(c context.Context, userID, status uint64) error
}

type user struct {
	factory services.FactoryImp
}

func NewUser(f services.FactoryImp) IUSer {
	return &user{
		factory: f,
	}
}

func (u *user) Login(params *form.UserLoginForm) (user *sys.User) {
	return u.factory.User().Login(params)

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
	u.factory.User().Register(userInfo)
}

func (u *user) Update(c context.Context, userID uint64, params *form.UpdateUserFrom) (err error) {
	userInfo := &sys.User{
		UserName:    params.UserName,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}
	err = u.factory.User().Update(userID, userInfo)
	if err != nil {
		log.Logger.Error(err)
	}
	return nil
}

func (u *user) GetUserInfoById(c context.Context, id uint64) (user *sys.User, err error) {
	user, err = u.factory.User().GetUserInfoById(id)
	if err != nil {
		return nil, err
	}
	return
}

func (u *user) GetUserList(pageParam *model.PageParam) *form.PageUser {
	return u.factory.User().GetUserList(pageParam)

}

func (u *user) DeleteUserByUserId(c context.Context, id uint64) (err error) {

	if err = u.factory.User().DeleteUserByUserId(id); err != nil {
		log.Logger.Error(err)
		return
	}

	if err = u.factory.Authentication().DeleteUser(id); err != nil {
		log.Logger.Error(err)
		return
	}

	return
}

func (u *user) CheckUserExist(c context.Context, userName string) bool {
	return u.factory.User().CheckUserExist(userName)
}

func (u *user) GetByName(c context.Context, name string) (user *sys.User, err error) {
	user, err = u.factory.User().GetByName(name)
	if err != nil {
		log.Logger.Error(err)
	}
	return
}

// GetRoleIDByUser 查询用户角色
func (u *user) GetRoleIDByUser(c context.Context, userID uint64) (roles *[]sys.Role, err error) {
	roles, err = u.factory.User().GetRoleIdbyUser(userID)
	if err != nil {
		log.Logger.Error(err)
	}
	return
}

// SetUserRoles 分配用户角色
func (u *user) SetUserRoles(c context.Context, userID uint64, roleIDS []uint64) (err error) {
	// 添加规则到rules表
	err = u.factory.Authentication().AddRoleForUser(userID, roleIDS)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	// 配置role_users表
	err = u.factory.User().SetUserRoles(userID, roleIDS)
	if err != nil { // 如果失败,则清除rules已添加的规则
		log.Logger.Error(err)
		for _, roleId := range roleIDS {
			err = u.factory.Authentication().DeleteRoleWithUser(userID, roleId)
			if err != nil {
				log.Logger.Error(err)
				break
			}
		}
		return
	}
	return
}

// GetButtonsByUserID 获取菜单按钮
func (u *user) GetButtonsByUserID(userID uint64) *[]string {
	menus := u.factory.User().GetButtonsByUserID(userID)

	var res []string
	for _, v := range *menus {
		if v.Code != "" {
			res = append(res, v.Code)
		}
	}

	return &res
}

// GetLeftMenusByUserID 根据用户ID获取左侧菜单
func (u *user) GetLeftMenusByUserID(userID uint64) (menus *[]sys.Menu) {
	return u.factory.User().GetLeftMenusByUserID(userID)

}

func (u *user) UpdateStatus(c context.Context, userId, status uint64) (err error) {
	err = u.factory.User().UpdateStatus(userId, status)
	if err != nil {
		log.Logger.Error(err)
	}
	return
}
