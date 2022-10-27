package services

//
//func Login(params form.UserLoginForm) (user *sys.User, err error) {
//	err = global.App.DB.Where("user_name = ?", params.UserName).First(&user).Error
//	if err != nil {
//		return
//	}
//	return
//}
//
//func Register(c *gin.Context, params form.RegisterUserForm) (err error, user sys.User) {
//
//	user = sys.User{
//		UserName: params.UserName,
//		Password: util.BcryptMake([]byte(params.Password)),
//		Mobile:   params.Mobile,
//		Email:    params.Email,
//	}
//	err = global.App.DB.Create(&user).Error
//	if err != nil {
//		return err, user
//	}
//	return nil, user
//}
//
//func GetUserInfoById(id uint64) (err error, user *sys.User) {
//	err = global.App.DB.First(&user, id).Error
//	if err != nil {
//		return
//	}
//	return
//}
//
//func GetUserList() (err error, user []sys.User) {
//	err = global.App.DB.Find(&user).Error
//	if err != nil {
//		return
//	}
//	return
//}
//
//func DeleteUserByUserId(id uint64) error {
//	return global.App.DB.Where("id = ?", id).Delete(&sys.User{}).Error
//}
//
//func CheckUserExist(userName string) bool {
//	err := global.App.DB.Where("user_name = ?", userName).First(&sys.User{}).Error
//	if err != nil {
//		return false
//	}
//
//	if err == gorm.ErrRecordNotFound {
//		return false
//	}
//	return true
//}
