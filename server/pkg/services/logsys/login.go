package logsys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type ILoginLog interface {
	Get(id uint64) *logsys.LogLogin
	List(query *model.PageParam, condition *logsys.LogLogin) *form.PageResult
	Add(*logsys.LogLogin)
	Delete(ids []uint64)
	DeleteAll()
}

type loginLog struct {
	db *gorm.DB
}

func (l *loginLog) Get(id uint64) (log *logsys.LogLogin) {
	restfulx.ErrIsNilRes(l.db.Where("id = ?", id).First(&log).Error, restfulx.GetResourceErr)
	return log
}

func (l *loginLog) List(query *model.PageParam, condition *logsys.LogLogin) (result *form.PageResult) {
	db := l.db
	logs := make([]logsys.LogLogin, 0)
	offset := (query.Page - 1) * query.Limit
	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.Username != "" {
		db = db.Where("username like ?", condition.Username)
	}

	restfulx.ErrIsNilRes(db.Model(&logsys.LogLogin{}).
		Count(&result.Total).
		Error,
		restfulx.GetResourceErr)

	restfulx.ErrIsNilRes(db.Model(&logsys.LogLogin{}).
		Offset(offset).
		Find(&logs).Error,
		restfulx.GetResourceErr)
	result.Data = logs

	return result
}

func (l *loginLog) Add(login *logsys.LogLogin) {
	restfulx.ErrIsNilRes(l.db.Create(login).Error, restfulx.OperatorErr)
}

func (l *loginLog) Delete(ids []uint64) {
	restfulx.ErrIsNilRes(l.db.Delete(&logsys.LogLogin{}).Where("id in (?)", ids).Error, restfulx.OperatorErr)
}

func (l *loginLog) DeleteAll() {
	restfulx.ErrIsNilRes(l.db.Exec("DELETE FROM log_login ").Error, restfulx.OperatorErr)
}

func NewLoginLog(db *gorm.DB) *loginLog {
	return &loginLog{db: db}
}
