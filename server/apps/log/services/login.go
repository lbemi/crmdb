package services

import (
	"github.com/lbemi/lbemi/apps/log/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type LoginLogGetter interface {
	Login() ILoginLog
}

type ILoginLog interface {
	Get(id uint64) *entity.LogLogin
	List(query *entity2.PageParam, condition *entity.LogLogin) *entity2.PageResult
	Add(*entity.LogLogin)
	Delete(ids []uint64)
	DeleteAll()
}

type LoginLog struct {
	db *gorm.DB
}

func (l *LoginLog) Get(id uint64) *entity.LogLogin {
	log := &entity.LogLogin{}
	restfulx.ErrNotNilDebug(l.db.Where("id = ?", id).First(&log).Error, restfulx.GetResourceErr)
	return log
}

func (l *LoginLog) List(query *entity2.PageParam, condition *entity.LogLogin) *entity2.PageResult {
	result := &entity2.PageResult{}
	db := l.db
	logs := make([]*entity.LogLogin, 0)
	offset := (query.Page - 1) * query.Limit
	if condition.Status != "" {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.Username != "" {
		db = db.Where("username like ?", "%"+condition.Username+"%")
	}

	restfulx.ErrNotNilDebug(db.Model(&entity.LogLogin{}).
		Count(&result.Total).
		Error,
		restfulx.GetResourceErr)

	restfulx.ErrNotNilDebug(db.Model(&entity.LogLogin{}).
		Order("loginTime DESC").
		Offset(offset).
		Limit(query.Limit).
		Find(&logs).Error,
		restfulx.GetResourceErr)

	result.Data = logs

	return result
}

func (l *LoginLog) Add(logLogin *entity.LogLogin) {
	restfulx.ErrNotNilDebug(l.db.Create(logLogin).Error, restfulx.OperatorErr)
}

func (l *LoginLog) Delete(ids []uint64) {
	restfulx.ErrNotNilDebug(l.db.Where("id in (?)", ids).Delete(&entity.LogLogin{}).Error, restfulx.OperatorErr)
}

func (l *LoginLog) DeleteAll() {
	restfulx.ErrNotNilDebug(l.db.Exec("DELETE FROM log_login ").Error, restfulx.OperatorErr)
}

func NewLogin(db *gorm.DB) ILoginLog {
	return &LoginLog{db: db}
}
