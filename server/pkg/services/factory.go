package services

import (
	"github.com/casbin/casbin/v2"
	"github.com/lbemi/lbemi/pkg/services/auth"
	"github.com/lbemi/lbemi/pkg/services/sys"
	"gorm.io/gorm"
)

type IDbFactory interface {
	Authentication() auth.AuthenticationInterface
	User() sys.IUSer
	Role() sys.IRole
	Menu() sys.IMenu
}

type DbFactory struct {
	db      *gorm.DB
	enforce *casbin.Enforcer
}

func (f *DbFactory) Authentication() auth.AuthenticationInterface {
	return auth.NewAuthentication(f.db, f.enforce)
}

func (f *DbFactory) User() sys.IUSer {
	return sys.NewUser(f.db)
}

func (f *DbFactory) Role() sys.IRole {
	return sys.NewRole(f.db)
}

func (f *DbFactory) Menu() sys.IMenu {
	return sys.NewMenu(f.db)
}

func NewDbFactory(db *gorm.DB, enforcer *casbin.Enforcer) IDbFactory {
	return &DbFactory{
		db:      db,
		enforce: enforcer,
	}
}
