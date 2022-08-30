package bootstrap

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lbemi/lbemi/pkg/global"
	"github.com/lbemi/lbemi/pkg/model/cas"
)

func InitCasbinEnforcer() *casbin.Enforcer {
	rbacRules :=
		`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
	m, _ := model.NewModelFromString(rbacRules)
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(global.App.DB, &cas.CasbinModel{})

	enforcer, _ := casbin.NewEnforcer(m, adapter)
	_ = enforcer.LoadPolicy()
	enforcer.SavePolicy()
	return enforcer
}
