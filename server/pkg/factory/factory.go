package factory

import (
	"github.com/lbemi/lbemi/pkg/services/sys/user"
	"gorm.io/gorm"
)

type IDbFactory interface {
	User() user.IUSer
}

type DbFactory struct {
	db *gorm.DB
}

func (f *DbFactory) User() user.IUSer {
	return user.NewUser(f.db)
}

func NewDbFactory(db *gorm.DB) IDbFactory {
	return &DbFactory{
		db: db,
	}
}
