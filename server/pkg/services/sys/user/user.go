package user

import (
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"gorm.io/gorm"
)

type IUSer interface {
	Login(params *form.UserLoginForm) (user *sys.User, err error)
	Register(params *sys.User) (err error)
	GetUserInfoById(id uint64) (user *sys.User, err error)
	GetUserList() (err error, user *[]sys.User)
	DeleteUserByUserId(id uint64) error
	CheckUserExist(userName string) bool
}

type user struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) IUSer {
	return &user{DB: db}
}

func (u *user) Login(params *form.UserLoginForm) (user *sys.User, err error) {
	err = u.DB.Where("user_name = ?", params.UserName).First(&user).Error
	if err != nil {
		return
	}
	return
}

func (u *user) Register(params *sys.User) (err error) {

	err = u.DB.Create(&params).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *user) GetUserInfoById(id uint64) (user *sys.User, err error) {
	err = u.DB.First(&user, id).Error
	if err != nil {
		return
	}
	return
}

func (u *user) GetUserList() (err error, user *[]sys.User) {
	err = u.DB.Find(&user).Error
	if err != nil {
		return
	}
	return
}

func (u *user) DeleteUserByUserId(id uint64) error {
	return u.DB.Where("id = ?", id).Delete(&sys.User{}).Error
}

func (u *user) CheckUserExist(userName string) bool {
	err := u.DB.Where("user_name = ?", userName).First(&sys.User{}).Error
	if err != nil {
		return false
	}

	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
