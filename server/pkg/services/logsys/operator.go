package logsys

//
//import (
//	"github.com/lbemi/lbemi/pkg/model"
//	"github.com/lbemi/lbemi/pkg/model/form"
//	"github.com/lbemi/lbemi/pkg/model/logsys"
//	"github.com/lbemi/lbemi/pkg/restfulx"
//
//	"gorm.io/gorm"
//)
//
//type IOperatorLog interface {
//	Get(id uint64) *logsys.LogOperator
//	List(query *model.PageParam, condition *logsys.LogOperator) *form.PageResult
//	Add(*logsys.LogOperator)
//	Delete(ids []uint64)
//	DeleteAll()
//}
//
//type operatorLog struct {
//	db *gorm.DB
//}
//
//func (l *operatorLog) Get(id uint64) (log *logsys.LogOperator) {
//	log = &logsys.LogOperator{}
//	restfulx.ErrNotNilDebug(l.db.Where("id = ?", id).First(&log).Error, restfulx.GetResourceErr)
//	return log
//}
//
//func (l *operatorLog) List(query *model.PageParam, condition *logsys.LogOperator) (result *form.PageResult) {
//	result = &form.PageResult{}
//	db := l.db
//	logs := make([]logsys.LogOperator, 0)
//	offset := (query.Page - 1) * query.Limit
//
//	if condition.BusinessType != "" {
//		db = db.Where("businessType = ?", condition.BusinessType)
//	}
//
//	if condition.Title != "" {
//		db = db.Where("title like ?", "%"+condition.Title+"%")
//	}
//
//	if condition.Name != "" {
//		db = db.Where("name like ?", "%"+condition.Name+"%")
//	}
//
//	if condition.Status == 200 {
//		db = db.Where("status = 200")
//	}
//	if condition.Status == 404 {
//		db = db.Where("status != 200")
//	}
//
//	restfulx.ErrNotNilDebug(db.Model(&logsys.LogOperator{}).
//		Count(&result.Total).
//		Error,
//		restfulx.GetResourceErr)
//
//	restfulx.ErrNotNilDebug(db.Model(&logsys.LogOperator{}).
//		Order("updated_at DESC").
//		Offset(offset).
//		Limit(query.Limit).
//		Find(&logs).Error,
//		restfulx.GetResourceErr)
//
//	result.Data = logs
//
//	return result
//}
//
//func (l *operatorLog) Add(ol *logsys.LogOperator) {
//	restfulx.ErrNotNilDebug(l.db.Create(ol).Error, restfulx.OperatorErr)
//}
//
//func (l *operatorLog) Delete(ids []uint64) {
//	restfulx.ErrNotNilDebug(l.db.Where("id in (?)", ids).Delete(&logsys.LogOperator{}).Error, restfulx.OperatorErr)
//}
//
//func (l *operatorLog) DeleteAll() {
//	restfulx.ErrNotNilDebug(l.db.Exec("DELETE FROM log_operator ").Error, restfulx.OperatorErr)
//}
//
//func NewOperatorLog(db *gorm.DB) *operatorLog {
//	return &operatorLog{db: db}
//}
