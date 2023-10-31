package logsys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/services"
)

type OperatorLogGetter interface {
	Operator() OperatorLogImp
}

type OperatorLogImp interface {
	Get(id uint64) *logsys.LogOperator
	List(query *model.PageParam, condition *logsys.LogOperator) *form.PageResult
	Add(*logsys.LogOperator)
	Delete(ids []uint64)
	DeleteAll()
}

type operator struct {
	factory services.Interface
}

func (l *operator) Get(id uint64) *logsys.LogOperator {
	return l.factory.Operator().Get(id)
}

func (l *operator) List(query *model.PageParam, condition *logsys.LogOperator) *form.PageResult {
	return l.factory.Operator().List(query, condition)
}

func (l *operator) Add(logLogin *logsys.LogOperator) {
	l.factory.Operator().Add(logLogin)
}

func (l *operator) Delete(ids []uint64) {
	l.factory.Operator().Delete(ids)
}

func (l *operator) DeleteAll() {
	l.factory.Operator().DeleteAll()
}

func NewOperator(f services.Interface) *operator {
	return &operator{factory: f}
}
