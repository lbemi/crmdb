package logsys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/services"
)

type LoginLogGetter interface {
	Login() LoginLogImp
}

type LoginLogImp interface {
	Get(id uint64) *logsys.LogLogin
	List(query *model.PageParam, condition *logsys.LogLogin) *form.PageResult
	Add(*logsys.LogLogin)
	Delete(ids []uint64)
	DeleteAll()
}

type login struct {
	factory services.FactoryImp
}

func (l *login) Get(id uint64) *logsys.LogLogin {
	return l.factory.Log().Get(id)
}

func (l *login) List(query *model.PageParam, condition *logsys.LogLogin) *form.PageResult {
	return l.factory.Log().List(query, condition)
}

func (l *login) Add(logLogin *logsys.LogLogin) {
	l.factory.Log().Add(logLogin)
}

func (l *login) Delete(ids []uint64) {
	l.factory.Log().Delete(ids)
}

func (l *login) DeleteAll() {
	l.factory.Log().DeleteAll()
}

func NewLogin(f services.FactoryImp) *login {
	return &login{factory: f}
}
