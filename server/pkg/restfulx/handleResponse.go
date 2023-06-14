package restfulx

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
)

const SuccessMsg = "success"

type Response struct {
	Code    int16       `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

//func (r *Response) ToJson() string {
//	marshalData, err := json.Marshal(r.Data)
//	if err != nil {
//		fmt.Println("marshal data to json failed")
//	}
//	return string(marshalData)
//}
//
//func (r *Response) IsSuccess() bool {
//	return r.Code == http.StatusOK
//}

func SuccessX(data interface{}) *Response {
	return &Response{Code: Success.code, Message: SuccessMsg, Data: data}
}

func SuccessRes(r *restful.Response, data interface{}) {
	r.WriteEntity(&Response{Code: Success.code, Message: SuccessMsg, Data: data})
}

func ErrorRes(r *restful.Response, err interface{}) {
	switch t := err.(type) {
	case *OpsError:
		r.WriteEntity(Error(t))
	case error:
		r.WriteEntity(ServerError())
		//log.Logger.Error(message)
	//case string:
	//	r.WriteEntity(ServerError())
	//	//log.Logger.Error(message)
	default:
		log.Logger.Error(err)
	}
}
