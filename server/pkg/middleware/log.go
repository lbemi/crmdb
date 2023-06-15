package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/core"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"

	"github.com/mssola/useragent"
)

type Fields map[string]interface{}

func LogHandler(rc *rctx.ReqCtx) error {
	// 记录日志
	go func() {
		defer func() {
			if r := recover(); r != nil {
				switch t := r.(type) {
				case *restfulx.OpsError:
					log.Logger.Error(t.Error())
				case error:
					log.Logger.Error(t)
				case string:
					log.Logger.Error(t)
				}
			}
		}()

		c := rc.Request
		// 请求操作不做记录
		if c.Request.Method == http.MethodGet || rc.LoginAccount == nil {
			return
		}
		if rc.RequirePermission == nil || !rc.RequirePermission.NeedToken {
			return
		}

		log := &logsys.LogOperator{
			Title:        rc.LogInfo.LogModule,
			BusinessType: "01",
			Method:       c.Request.Method,
			Name:         rc.LoginAccount.UserName,
			Url:          c.Request.URL.Path,
			Ip:           rc.ClientIP(),
			Status:       200,
		}
		if rc.Err != nil {
			switch t := rc.Err.(type) {
			case *restfulx.OpsError:
				log.Status = t.Code()
				log.ErrMsg = t.Error()
			case error:
				log.Status = 500
				log.ErrMsg = "服务器内部错误"
			case string:
				log.Status = 500
				log.ErrMsg = rc.Err.(string)
			}
		}
		if c.Request.Method == "POST" {
			log.BusinessType = "01"
		} else if c.Request.Method == "PUT" {
			log.BusinessType = "02"
		} else if c.Request.Method == "DELETE" {
			log.BusinessType = "03"
		} else {
			log.BusinessType = "04"
		}
		core.V1.Operator().Add(log)
	}()
	msg := getLogMsg(rc)

	if rc.Err != nil {
		// 如果是非自定义错误日志，则打印堆栈信息
		switch t := rc.Err.(type) {
		case *restfulx.OpsError:
			log.Logger.Error(msg, "|", t.Code(), " - "+t.Error())
		case error:
			log.Logger.Error(msg, t)
		case string:
			log.Logger.Error(msg, t)
		}
		return nil
	}

	log.Logger.Info(msg)
	return nil
}

func getLogMsg(rc *rctx.ReqCtx) string {

	req := rc.Request.Request
	ua := useragent.New(req.UserAgent())
	bName, bVersion := ua.Browser()

	msg := fmt.Sprintf("%v | %v | %dms | %v | %v query:(%v) | %v | %v  | %v-%v ",
		rc.Response.StatusCode(), req.Method, rc.Timed, rc.LogInfo.LogModule, req.URL.Path, req.URL.Query().Encode(),
		rc.ClientIP(), ua.OS(), bName, bVersion)
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
