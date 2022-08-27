package services

import (
	"errors"
	"fmt"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/util"
)

//type userService struct {
//}

func Register(params form.RegisterUserForm) (err error, user sys.User) {
	rs := global.App.DB.Where("user_name = ?", params.UserName).First(&sys.User{})
	if rs.RowsAffected != 0 {
		err = errors.New("用户已存在")
		return
	}
	user = sys.User{
		UserName: params.UserName,
		Password: util.BcryptMake([]byte(params.Password)),
		Mobile:   params.Mobile,
		Email:    params.Email,
	}
	err = global.App.DB.Create(&user).Error
	if err != nil {
		return err, user
		fmt.Println(err)
	}
	return nil, user
}

func Login(params form.UserLoginForm) (user *sys.User, err error) {
	err = global.App.DB.Where("user_name = ?", params.UserName).First(&user).Error
	if err != nil || !util.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户不存在或密码错误")
	}
	return
}
