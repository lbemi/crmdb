package logsys

import (
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type LoginLogGetter interface {
	Login() ILoginLog
}

type ILoginLog interface {
	Get(id uint64) *logsys.LogLogin
	List(query *model.PageParam, condition *logsys.LogLogin) *form.PageResult
	Add(*logsys.LogLogin)
	Delete(ids []uint64)
	DeleteAll()
}

type LoginLog struct {
	db *gorm.DB
}

func (l *LoginLog) Get(id uint64) *logsys.LogLogin {
	log := &logsys.LogLogin{}
	restfulx.ErrNotNilDebug(l.db.Where("id = ?", id).First(&log).Error, restfulx.GetResourceErr)
	return log
}

func (l *LoginLog) List(query *model.PageParam, condition *logsys.LogLogin) *form.PageResult {
	result := &form.PageResult{}
	db := l.db
	logs := make([]*logsys.LogLogin, 0)
	offset := (query.Page - 1) * query.Limit
	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.Username != "" {
		db = db.Where("username like ?", "%"+condition.Username+"%")
	}

	restfulx.ErrNotNilDebug(db.Model(&logsys.LogLogin{}).
		Count(&result.Total).
		Error,
		restfulx.GetResourceErr)

	restfulx.ErrNotNilDebug(db.Model(&logsys.LogLogin{}).
		Order("loginTime DESC").
		Offset(offset).
		Limit(query.Limit).
		Find(&logs).Error,
		restfulx.GetResourceErr)

	result.Data = logs

	return result
}

func (l *LoginLog) Add(logLogin *logsys.LogLogin) {
	restfulx.ErrNotNilDebug(l.db.Create(logLogin).Error, restfulx.OperatorErr)
}

func (l *LoginLog) Delete(ids []uint64) {
	restfulx.ErrNotNilDebug(l.db.Where("id in (?)", ids).Delete(&logsys.LogLogin{}).Error, restfulx.OperatorErr)
}

func (l *LoginLog) DeleteAll() {
	restfulx.ErrNotNilDebug(l.db.Exec("DELETE FROM log_login ").Error, restfulx.OperatorErr)
}

func NewLogin(db *gorm.DB) ILoginLog {
	return &LoginLog{db: db}
}
