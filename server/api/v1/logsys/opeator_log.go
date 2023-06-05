package logsys

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
)

func GetOperatorLog(rc *rctx.ReqCtx) {
	id := rctx.PathParamUint64(rc, "id")
	rc.ResData = core.V1.Operator().Get(id)
}

func ListOperatorLog(rc *rctx.ReqCtx) {
	query := model.PageParam{}
	query.Page = rctx.QueryDefaultInt(rc, "page", 0)
	query.Limit = rctx.QueryDefaultInt(rc, "limit", 10)
	condition := logsys.LogOperator{}
	condition.BusinessType = rctx.QueryParam(rc, "type")
	condition.Title = rctx.QueryParam(rc, "title")
	condition.Name = rctx.QueryParam(rc, "name")
	status := rctx.QueryParam(rc, "status")
	if status == "normal" {
		condition.Status = 200
	}
	if status == "failed" {
		condition.Status = 404
	}
	rc.ResData = core.V1.Operator().List(&query, &condition)
}

func DeleteOperatorLogs(rc *rctx.ReqCtx) {
	ids := util.ParseStrInt64(rctx.QueryParam(rc, "ids"))
	core.V1.Operator().Delete(ids)
}

func DeleteAllOperatorLog(rc *rctx.ReqCtx) {
	core.V1.Operator().DeleteAll()
}
