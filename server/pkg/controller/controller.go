package controller

import (
	"github.com/lbemi/lbemi/pkg/controller/sys/user"
	"github.com/lbemi/lbemi/pkg/factory"
	"github.com/lbemi/lbemi/pkg/model/configs"
)

type IController interface {
	user.UserGetter
}

type Controller struct {
	Config    *configs.Config
	DbFactory factory.DbFactory
}

func NewController(factory factory.DbFactory) IController {
	return &Controller{
		DbFactory: factory,
	}
}

func (c *Controller) User() user.IUSer {
	return user.NewUser(c)
}
