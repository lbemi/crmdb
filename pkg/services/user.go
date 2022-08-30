package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/util"
	"strconv"
)

func Login(params form.UserLoginForm) (user *sys.User, err error) {
	err = global.App.DB.Where("user_name = ?", params.UserName).First(&user).Error
	if err != nil {
		return
	}
	return
}

func Register(c *gin.Context, params form.RegisterUserForm) (err error, user sys.User) {
	rs := global.App.DB.Where("user_name = ?", params.UserName).First(&sys.User{})
	if rs.RowsAffected != 0 {
		err = errors.New("用户已存在")
		return
	}
	//adminId, ok := c.Get("id")
	//var admin uint64
	//if ok {
	//	admin, err = strconv.ParseUint(adminId.(string), 10, 64)
	//}

	user = sys.User{
		UserName: params.UserName,
		Password: util.BcryptMake([]byte(params.Password)),
		Mobile:   params.Mobile,
		Email:    params.Email,
	}
	err = global.App.DB.Create(&user).Error
	if err != nil {
		return err, user
	}
	return nil, user
}

func GetUserInfoById(id string) (err error, user sys.User) {
	intId, err := strconv.Atoi(id)
	err = global.App.DB.First(&user, intId).Error
	if err != nil {
		return
	}
	return
}

func GetUserInfos(id string) (err error, user []sys.User) {
	err = global.App.DB.Find(&user).Error
	if err != nil {
		return
	}
	return
}
