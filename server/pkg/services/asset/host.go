package asset

import (
	"context"
	"github.com/lbemi/lbemi/pkg/restfulx"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"gorm.io/gorm"
)

type host struct {
	db *gorm.DB
}

func NewHost(DB *gorm.DB) IHost {
	return &host{db: DB}
}

type IHost interface {
	Create(ctx context.Context, host *asset.Host)
	Delete(ctx context.Context, hostId int64)
	Update(ctx context.Context, hostId int64, host *asset.Host)
	List(page, limit int) *form.PageHost
	GetByHostId(ctx context.Context, hostId int64) (host *asset.Host)
	UpdateFiledStatus(ctx context.Context, hostId int64, updateFiled string, status int8)
}

func (m *host) Create(ctx context.Context, host *asset.Host) {
	restfulx.ErrNotNilDebug(m.db.Create(host).Error, restfulx.OperatorErr)
}

func (m *host) Delete(ctx context.Context, hostId int64) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Delete(&asset.Host{}).Error, restfulx.OperatorErr)
}

func (m *host) Update(ctx context.Context, hostId int64, host *asset.Host) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Updates(host).Error, restfulx.OperatorErr)
}

func (m *host) List(page, limit int) *form.PageHost {
	var (
		hostList []asset.Host
		total    int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(m.db.Find(&hostList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(m.db.Model(&asset.Host{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageHost{
			Hosts: hostList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(m.db.Limit(limit).Offset((page-1)*limit).
		Find(&hostList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(m.db.Model(&asset.Host{}).Count(&total).Error, restfulx.OperatorErr)

	res := &form.PageHost{
		Hosts: hostList,
		Total: total,
	}
	return res
}

func (m *host) GetByHostId(ctx context.Context, hostId int64) (host *asset.Host) {
	host = &asset.Host{}
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Find(&host).Error, restfulx.OperatorErr)
	return host
}

func (m *host) UpdateFiledStatus(ctx context.Context, hostId int64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}
