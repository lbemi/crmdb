package cmd

import (
	"github.com/lbemi/lbemi/pkg/bootstrap"
	"github.com/lbemi/lbemi/pkg/cache"
	"github.com/lbemi/lbemi/pkg/services"
)

func TestInit(c string) services.FactoryImp {
	config := bootstrap.InitializeConfig(c)
	db := bootstrap.InitializeDB(config)
	// 初始化casbin enforcer
	//enforcer := bootstrap.InitPolicyEnforcer(db)
	clientStore := cache.NewClientStore()
	// 初始化dbFactory
	return services.NewDbFactory(db, nil, clientStore)
}
