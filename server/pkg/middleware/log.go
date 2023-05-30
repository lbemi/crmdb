package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/util"
	"github.com/mssola/useragent"
	"reflect"
)

type Fields map[string]interface{}

func LogHandler(rc *rctx.ReqCtx) error {

	if err := rc.Err; err != nil {
		log.Logger.Error(getLogMsg(rc), err)
		return nil
	}
	log.Logger.Info(getLogMsg(rc))
	return nil
}

func getLogMsg(rc *rctx.ReqCtx) string {
	req := rc.Request.Request
	ua := useragent.New(req.UserAgent())
	bName, bVersion := ua.Browser()

	msg := fmt.Sprintf("%v | %v | %dms | %v | %v query:(%v) | %v | %v  | %v-%v ",
		rc.Response.StatusCode(), req.Method, rc.Timed, rc.LogInfo.LogModule, req.URL.Path, req.URL.Query().Encode(),
		req.RemoteAddr, ua.OS(), bName, bVersion)
	if rc.LoginAccount != nil {
		msg += fmt.Sprintf("| user: %v ",
			rc.LoginAccount.UserName)
	}

	// 返回结果不为空，则记录返回结果
	if rc.LogInfo.LogRes && !util.IsBlank(reflect.ValueOf(rc.ResData)) {
		respB, _ := json.Marshal(rc.ResData)
		msg = msg + fmt.Sprintf("\n<-- %s", string(respB))
	}

	return msg
}
