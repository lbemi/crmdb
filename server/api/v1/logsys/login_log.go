package logsys

import (
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
)

func GetLoginLog(rc *rctx.ReqCtx) {
	id := rc.PathParamUint64("id")
	rc.ResData = core.V1.Login().Get(id)
}

func ListLoginLog(rc *rctx.ReqCtx) {
	query := &model.PageParam{}
	query.Page = rc.QueryDefaultInt("page", 0)
	query.Limit = rc.QueryDefaultInt("limit", 10)
	condition := &logsys.LogLogin{}
	condition.Status = rc.Query("status")
	condition.Username = rc.Query("name")
	rc.ResData = core.V1.Login().List(query, condition)
}

func DeleteLoginLogs(rc *rctx.ReqCtx) {
	ids := util.ParseStrInt64(rc.Query("ids"))
	core.V1.Login().Delete(ids)
}

func DeleteAllLoginLog(rc *rctx.ReqCtx) {
	core.V1.Login().DeleteAll()
}