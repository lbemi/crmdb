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
	Delete(ctx context.Context, hostId uint64)
	Update(ctx context.Context, hostId uint64, host *asset.Host)
	List(ctx context.Context, page, limit int, groups []uint64, ip, label, description string) *form.PageHost
	GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host)
	GetByGroup(ctx context.Context, groups []uint64, page, limit int) *form.PageResult
	UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8)
}

func (m *host) Create(ctx context.Context, host *asset.Host) {
	restfulx.ErrNotNilDebug(m.db.Create(host).Error, restfulx.OperatorErr)
}

func (m *host) Delete(ctx context.Context, hostId uint64) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Delete(&asset.Host{}).Error, restfulx.OperatorErr)
}

func (m *host) Update(ctx context.Context, hostId uint64, host *asset.Host) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Updates(host).Error, restfulx.OperatorErr)
}

func (m *host) List(ctx context.Context, page, limit int, groups []uint64, ip, label, description string) *form.PageHost {
	var (
		hostList []asset.Host
		total    int64
	)
	db := m.db

	if len(groups) > 0 {
		db = db.Where("group_id in (?)", groups)
	}
	if ip != "" {
		db = db.Where("ip LIKE ?", "%"+ip+"%")
	}
	if label != "" {
		db = db.Where("labels LIKE ?", "%"+label+"%")
	}
	if description != "" {
		db = db.Where("remark LIKE ?", "%"+description+"%")
	}

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(db.Find(&hostList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(db.Model(&asset.Host{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageHost{
			Hosts: hostList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(db.Limit(limit).Offset((page-1)*limit).
		Find(&hostList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(db.Model(&asset.Host{}).Count(&total).Error, restfulx.OperatorErr)

	res := &form.PageHost{
		Hosts: hostList,
		Total: total,
	}
	return res
}

func (m *host) GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host) {
	host = &asset.Host{}
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Find(&host).Error, restfulx.OperatorErr)
	return host
}

func (m *host) GetByGroup(ctx context.Context, groups []uint64, page, limit int) *form.PageResult {
	var (
		hostList []*asset.Host
		total    int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(m.db.Where("group_id in (?)", groups).Find(&hostList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(m.db.Where("group_id in (?)", groups).Model(&asset.Host{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageResult{
			Data:  hostList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(m.db.Where("group_id in (?)", groups).Limit(limit).Offset((page-1)*limit).
		Find(&hostList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(m.db.Model(&asset.Host{}).Where("group_id in (?)", groups).Count(&total).Error, restfulx.OperatorErr)

	res := &form.PageResult{
		Data:  hostList,
		Total: total,
	}
	return res
}

func (m *host) UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", hostId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}
