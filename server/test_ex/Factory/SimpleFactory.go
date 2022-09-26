package Factory

import "github.com/lbemi/lbemi/test_ex/models"

const (
	FrontUser = iota
	AdminUser
)

type UserType int

func CreateUser(UserType int) models.UserCreateFunc {
	switch UserType {
	case FrontUser:
		return models.NewUser()
	case AdminUser:
		return models.NewAdmin()
	default:
		return models.NewUser()
	}
}
