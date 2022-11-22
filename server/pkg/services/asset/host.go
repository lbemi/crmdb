package asset

import (
	"context"
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
	Create(ctx context.Context, host *asset.Host) error
	Delete(ctx context.Context, hostId int64) error
	Update(ctx context.Context, hostId int64, host *asset.Host) error
	List(page, limit int) (*form.PageHost, error)
	GetByHostId(ctx context.Context, hostId int64) (host *asset.Host, err error)
	UpdateFiledStatus(ctx context.Context, hostId int64, updateFiled string, status int8) error
}

func (m *host) Create(ctx context.Context, host *asset.Host) error {
	return m.db.Create(host).Error
}

func (m *host) Delete(ctx context.Context, hostId int64) error {
	return m.db.Where("id = ?", hostId).Delete(&asset.Host{}).Error
}

func (m *host) Update(ctx context.Context, hostId int64, host *asset.Host) error {
	return m.db.Where("id = ?", hostId).Updates(host).Error
}

func (m *host) List(page, limit int) (*form.PageHost, error) {
	var (
		hostList []asset.Host
		total    int64
		err      error
	)

	// 全量查询
	if page == 0 && limit == 0 {
		if tx := m.db.Find(&hostList); tx.Error != nil {
			return nil, tx.Error
		}

		if err := m.db.Model(&asset.Host{}).Count(&total).Error; err != nil {
			return nil, err
		}

		res := &form.PageHost{
			Hosts: hostList,
			Total: total,
		}
		return res, err
	}

	//分页数据
	if err := m.db.Limit(limit).Offset((page - 1) * limit).
		Find(&hostList).Error; err != nil {
		return nil, err
	}

	if err := m.db.Model(&asset.Host{}).Count(&total).Error; err != nil {
		return nil, err
	}

	res := &form.PageHost{
		Hosts: hostList,
		Total: total,
	}
	return res, err
}

func (m *host) GetByHostId(ctx context.Context, hostId int64) (host *asset.Host, err error) {
	err = m.db.Where("id = ?", hostId).Find(&host).Error
	return
}

func (m *host) UpdateFiledStatus(ctx context.Context, hostId int64, updateFiled string, status int8) error {
	return m.db.Where("id = ?", hostId).Update(updateFiled, status).Error
}
