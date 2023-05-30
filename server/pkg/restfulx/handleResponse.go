package restfulx

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
)

func SuccessRes(r *restful.Response, data interface{}) {
	r.WriteEntity(SuccessX(data))
}

func ErrorRes(r *restful.Response, err interface{}) {
	switch t := err.(type) {
	case OpsError:
		r.WriteEntity(Error(t))
	case error:
		r.WriteEntity(ServerError())
		//log.Logger.Error(message)
	case string:
		r.WriteEntity(ServerError())
		//log.Logger.Error(message)
	default:
		log.Logger.Error(err)
	}
}
