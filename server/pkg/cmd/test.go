package cmd

import (
	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/common/store"
	"github.com/lbemi/lbemi/pkg/services"
)

func TestInit(c string) services.Interface {
	config := bootstrap.InitializeConfig(c)
	db := bootstrap.InitializeDB(config)
	// 初始化casbin enforcer
	//enforcer := bootstrap.InitPolicyEnforcer(db)
	clientStore := store.NewClientStore()
	// 初始化dbFactory
	return services.NewDbFactory(db, nil, clientStore)
}
