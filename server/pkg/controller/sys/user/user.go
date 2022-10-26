package user

import (
	"context"
	"github.com/lbemi/lbemi/pkg/controller"
	"github.com/lbemi/lbemi/pkg/factory"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/util"
)

type UserGetter interface {
	User() IUSer
}

type IUSer interface {
	Login(c context.Context, params *form.UserLoginForm) (user *sys.User, err error)
	Register(c context.Context, params *form.RegisterUserForm) (err error)
	GetUserInfoById(c context.Context, id uint64) (user *sys.User, err error)
	GetUserList(c context.Context) (err error, user *[]sys.User)
	DeleteUserByUserId(c context.Context, id uint64) error
	CheckUserExist(c context.Context, userName string) bool
}

type user struct {
	factory factory.DbFactory
}

func NewUser(c *controller.Controller) IUSer {
	return &user{
		factory: c.DbFactory,
	}
}

func (u *user) Login(c context.Context, params *form.UserLoginForm) (user *sys.User, err error) {
	user, err = u.factory.User().Login(params)
	if err != nil {
		return
	}
	return
}

func (u *user) Register(c context.Context, params *form.RegisterUserForm) (err error) {

	user := &sys.User{
		UserName: params.UserName,
		Password: util.BcryptMake([]byte(params.Password)),
		Mobile:   params.Mobile,
		Email:    params.Email,
	}
	err = u.factory.User().Register(user)
	if err != nil {
		return err
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

func (u *user) GetUserList(c context.Context) (err error, user *[]sys.User) {
	err, user = u.factory.User().GetUserList()
	if err != nil {
		return
	}
	return
}

func (u *user) DeleteUserByUserId(c context.Context, id uint64) error {
	return u.factory.User().DeleteUserByUserId(id)
}

func (u *user) CheckUserExist(c context.Context, userName string) bool {
	return u.factory.User().CheckUserExist(userName)
}
