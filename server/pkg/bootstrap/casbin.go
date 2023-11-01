package bootstrap

import (
	"fmt"
	"github.com/lbemi/lbemi/pkg/config"

	"github.com/casbin/casbin/v2"
	csmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// InitPolicyEnforcer TODO: 整体优化
func InitPolicyEnforcer(db *gorm.DB) (enforcer *casbin.SyncedEnforcer) {
	rbacRules :=
		`
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "21220821"
	`
	// 加载鉴权规则
	m, err := csmodel.NewModelFromString(rbacRules)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 调用gorm创建casbin_rule表
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &config.Rule{}, "rules")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建鉴权器enforcer（使用gorm adapter）
	enforcer, err = casbin.NewSyncedEnforcer(m, adapter)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 	加载权限
	err = enforcer.LoadPolicy()
	if err != nil {
		fmt.Println(err)
		return
	}

	return enforcer

}
