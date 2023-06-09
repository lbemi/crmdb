package logsys

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
)

func GetOperatorLog(rc *rctx.ReqCtx) {
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Operator().Get(id)
}

func ListOperatorLog(rc *rctx.ReqCtx) {
	query := model.PageParam{}
	query.Page = rc.QueryDefaultInt("page", 0)
	query.Limit = rc.QueryDefaultInt("limit", 10)
	condition := logsys.LogOperator{}
	condition.BusinessType = rc.QueryParam("type")
	condition.Title = rc.QueryParam("title")
	condition.Name = rc.QueryParam("name")
	status := rc.QueryParam("status")
	if status == "normal" {
		condition.Status = 200
	}
	if status == "failed" {
		condition.Status = 404
	}
	rc.ResData = core.V1.Operator().List(&query, &condition)
}

func DeleteOperatorLogs(rc *rctx.ReqCtx) {
	ids := util.ParseStrInt64(rc.QueryParam("ids"))
	core.V1.Operator().Delete(ids)
}

func DeleteAllOperatorLog(rc *rctx.ReqCtx) {
	core.V1.Operator().DeleteAll()
}
