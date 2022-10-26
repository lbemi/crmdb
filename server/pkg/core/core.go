package core

import (
	"github.com/lbemi/lbemi/pkg/sys/user"
	"gorm.io/gorm"
)

type ICore interface {
	User() user.IUSer
}

type Core struct {
	db *gorm.DB
}

func (c *Core) User() user.IUSer {
	return user.NewUser(c.db)
}
