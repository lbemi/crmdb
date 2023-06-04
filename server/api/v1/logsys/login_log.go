package logsys

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
)

func GetLoginLog(rc *rctx.ReqCtx) {
	id := rctx.PathParamUint64(rc, "id")
	rc.ResData = core.V1.Login().Get(id)
}

func ListLoginLog(rc *rctx.ReqCtx) {
	query := model.PageParam{}
	query.Page = rctx.QueryDefaultInt(rc, "page", 0)
	query.Limit = rctx.QueryDefaultInt(rc, "limit", 10)
	condition := logsys.LogLogin{}
	condition.Status = rctx.QueryParam(rc, "status")
	condition.Username = rctx.QueryParam(rc, "name")
	rc.ResData = core.V1.Login().List(&query, &condition)
}

func DeleteLoginLogs(rc *rctx.ReqCtx) {
	ids := util.ParseStrInt64(rctx.QueryParam(rc, "ids"))
	core.V1.Login().Delete(ids)
}

func DeleteAllLoginLog(rc *rctx.ReqCtx) {
	core.V1.Login().DeleteAll()
}
