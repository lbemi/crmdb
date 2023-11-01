package services

import (
	"github.com/lbemi/lbemi/apps/log/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type OperatorLogGetter interface {
	Operator() IOperatorLog
}

type IOperatorLog interface {
	Get(id uint64) *entity.LogOperator
	List(query *entity2.PageParam, condition *entity.LogOperator) *entity2.PageResult
	Add(*entity.LogOperator)
	Delete(ids []uint64)
	DeleteAll()
}

type OperatorLog struct {
	db *gorm.DB
}

func (l *OperatorLog) Get(id uint64) *entity.LogOperator {
	log := &entity.LogOperator{}
	restfulx.ErrNotNilDebug(l.db.Where("id = ?", id).First(&log).Error, restfulx.GetResourceErr)
	return log
}

func (l *OperatorLog) List(query *entity2.PageParam, condition *entity.LogOperator) *entity2.PageResult {
	result := &entity2.PageResult{}
	db := l.db
	logs := make([]entity.LogOperator, 0)
	offset := (query.Page - 1) * query.Limit

	if condition.BusinessType != "" {
		db = db.Where("businessType = ?", condition.BusinessType)
	}

	if condition.Title != "" {
		db = db.Where("title like ?", "%"+condition.Title+"%")
	}

	if condition.Name != "" {
		db = db.Where("name like ?", "%"+condition.Name+"%")
	}

	if condition.Status == 200 {
		db = db.Where("status = 200")
	}
	if condition.Status == 404 {
		db = db.Where("status != 200")
	}

	restfulx.ErrNotNilDebug(db.Model(&entity.LogOperator{}).
		Count(&result.Total).
		Error,
		restfulx.GetResourceErr)

	restfulx.ErrNotNilDebug(db.Model(&entity.LogOperator{}).
		Order("updated_at DESC").
		Offset(offset).
		Limit(query.Limit).
		Find(&logs).Error,
		restfulx.GetResourceErr)

	result.Data = logs

	return result
}

func (l *OperatorLog) Add(operatorLog *entity.LogOperator) {
	restfulx.ErrNotNilDebug(l.db.Create(operatorLog).Error, restfulx.OperatorErr)
}

func (l *OperatorLog) Delete(ids []uint64) {
	restfulx.ErrNotNilDebug(l.db.Where("id in (?)", ids).Delete(&entity.LogOperator{}).Error, restfulx.OperatorErr)
}

func (l *OperatorLog) DeleteAll() {
	restfulx.ErrNotNilDebug(l.db.Exec("DELETE FROM log_operator ").Error, restfulx.OperatorErr)
}

func NewOperator(db *gorm.DB) IOperatorLog {
	return &OperatorLog{db: db}
}
