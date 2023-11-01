package api

import (
	"github.com/lbemi/lbemi/apps/log/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
)

func GetOperatorLog(rc *rctx.ReqCtx) {
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Operator().Get(id)
}

func ListOperatorLog(rc *rctx.ReqCtx) {
	query := entity2.PageParam{}
	query.Page = rc.QueryDefaultInt("page", 0)
	query.Limit = rc.QueryDefaultInt("limit", 10)
	condition := entity.LogOperator{}
	condition.BusinessType = rc.Query("type")
	condition.Title = rc.Query("title")
	condition.Name = rc.Query("name")
	status := rc.Query("status")
	if status == "normal" {
		condition.Status = 200
	}
	if status == "failed" {
		condition.Status = 404
	}
	rc.ResData = core.V1.Operator().List(&query, &condition)
}

func DeleteOperatorLogs(rc *rctx.ReqCtx) {
	ids := util.ParseStrInt64(rc.Query("ids"))
	core.V1.Operator().Delete(ids)
}

func DeleteAllOperatorLog(rc *rctx.ReqCtx) {
	core.V1.Operator().DeleteAll()
}
