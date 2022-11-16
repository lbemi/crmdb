package asset

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"

	"gorm.io/gorm"
)

type HostGetter interface {
	Host() IHost
}

type host struct {
	factory services.IDbFactory
}

func NewHost(f services.IDbFactory) IHost {
	return &host{
		factory: f,
	}
}

// IHost 主机操作接口
type IHost interface {
	Create(ctx context.Context, host *asset.HostReq) error
	Delete(ctx context.Context, hostId int64) error
	Update(ctx context.Context, hostId int64, host *asset.HostReq) error
	List(ctx context.Context, page, limit int) (*form.PageHost, error)
	GetByHostId(ctx context.Context, hostId int64) (host *asset.Host, err error)
	UpdateFiledStatus(ctx context.Context, hostId int64, updateFiled string, status int8) error
	CheckHostExist(ctx context.Context, hostId int64) bool
}

func (m *host) Create(ctx context.Context, host *asset.HostReq) error {
	return m.factory.Host().Create(ctx, &asset.Host{
		Label:      host.Label,
		Remark:     host.Remark,
		Ip:         host.Ip,
		Port:       host.Port,
		Username:   host.Username,
		AuthMethod: host.AuthMethod,
		Password:   host.Password,
		Secret:     host.Secret,
		Status:     host.Status,
		EnableSSH:  host.EnableSSH,
	})
}

func (m *host) Delete(ctx context.Context, hostId int64) error {
	err := m.factory.Host().Delete(ctx, hostId)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (m *host) Update(ctx context.Context, hostId int64, host *asset.HostReq) error {
	err := m.factory.Host().Update(ctx, hostId, &asset.Host{
		Label:      host.Label,
		Remark:     host.Remark,
		Ip:         host.Ip,
		Port:       host.Port,
		Username:   host.Username,
		AuthMethod: host.AuthMethod,
		Password:   host.Password,
		Status:     host.Status,
		EnableSSH:  host.EnableSSH,
	})
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (m *host) List(ctx context.Context, page, limit int) (*form.PageHost, error) {

	hosts, err := m.factory.Host().List(page, limit)
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return hosts, nil
}

func (m *host) GetByHostId(ctx context.Context, hostId int64) (host *asset.Host, err error) {
	host, err = m.factory.Host().GetByHostId(ctx, hostId)
	if err != nil {
		log.Logger.Error(err)
	}
	return
}

func (m *host) UpdateFiledStatus(ctx context.Context, hostId int64, updateFiled string, status int8) error {
	err := m.factory.Host().UpdateFiledStatus(ctx, hostId, updateFiled, status)
	if err != nil {
		log.Logger.Error(err)
	}
	return err
}

func (m *host) CheckHostExist(ctx context.Context, hostId int64) bool {
	_, err := m.factory.Host().GetByHostId(ctx, hostId)
	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
