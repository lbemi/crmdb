package api

import (
	"github.com/lbemi/lbemi/apps/asset/entity"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
)

func AddAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	account := &entity.Account{}
	rc.ShouldBind(account)
	core.V1.Account().Create(c, account)
}

func ListAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	page := rc.GetPageQueryParam()
	name := rc.Query("name")
	userName := rc.Query("userName")
	rc.ResData = core.V1.Account().List(c, page.Page, page.Limit, name, userName)
}

func GetAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Account().GetByAccountId(c, id)
}

func DeleteAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	id := rc.PathParamUint64("id")
	core.V1.Account().Delete(c, id)
}

func UpdateAccount(rc *rctx.ReqCtx) {
	c := rc.Request.Request.Context()
	account := &entity.Account{}
	rc.ShouldBind(account)
	core.V1.Account().Update(c, account.ID, account)
}
