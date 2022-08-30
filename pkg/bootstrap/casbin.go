package bootstrap

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/cas"
)

func InitCasbinEnforcer() *casbin.Enforcer {
	fmt.Println("///////////")
	rbacRules :=
		`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act)
`
	m, err := model.NewModelFromString(rbacRules)
	if err != nil {
		global.App.Log.Error(err.Error())
		return nil
	}
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.App.DB, &cas.CasbinModel{})
	if err != nil {
		global.App.Log.Error(err.Error())
		return nil
	}
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		global.App.Log.Error(err.Error())
		return nil
	}
	fmt.Println("******************")
	err = enforcer.LoadPolicy()
	if err != nil {
		global.App.Log.Error(err.Error())
		return nil
	}
	//enforcer.SavePolicy()
	fmt.Println("初始化。。。。Casbin")
	return enforcer
}
