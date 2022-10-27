package bootstrap

import (
	"github.com/casbin/casbin/v2"
)

func InitCasbinEnforcer() *casbin.Enforcer {
	//rbacRules :=
	//	`
	//[request_definition]
	//r = sub, obj, act
	//
	//[policy_definition]
	//p = sub, obj, act
	//
	//[role_definition]
	//g = _, _
	//
	//[policy_effect]
	//e = some(where (p.eft == allow))
	//
	//[matchers]
	//m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "1"
	//`
	//m, err := model.NewModelFromString(rbacRules)
	//if err != nil {
	//	global.App.Log.Error(err.Error())
	//	return nil
	//}
	//adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.App.DB, &cas.CasbinModel{})
	//
	//if err != nil {
	//	global.App.Log.Error(err.Error())
	//	return nil
	//}
	//enforcer, err := casbin.NewEnforcer(m, adapter)
	//
	//if err != nil {
	//	global.App.Log.Error(err.Error())
	//	return nil
	//}
	//err = enforcer.LoadPolicy()
	//if err != nil {
	//	global.App.Log.Error(err.Error())
	//	return nil
	//}
	////enforcer.SavePolicy()
	return nil
}
