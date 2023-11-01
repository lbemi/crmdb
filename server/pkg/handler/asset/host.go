package asset

import (
	"context"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
)

type HostGetter interface {
	Host() IHost
}

type host struct {
	db      *gorm.DB
	rbac    IResourceBindAccount
	account IAccount
}

func NewHost(db *gorm.DB, rbac IResourceBindAccount, account IAccount) IHost {
	return &host{
		db:      db,
		rbac:    rbac,
		account: account,
	}
}

// IHost 主机操作接口
type IHost interface {
	Create(ctx context.Context, host *asset.Host)
	Delete(ctx context.Context, hostId uint64)
	Update(ctx context.Context, hostId uint64, host *asset.Host)
	List(ctx context.Context, page, limit int, groups []uint64, ip, label, description string) *form.PageHost
	GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host)
	GetHostAccounts(ctx context.Context, hostId uint64) []*asset.Account
	GetByGroup(ctx context.Context, groups []uint64, page, limit int) *form.PageResult
	UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8)
	CheckHostExist(ctx context.Context, hostId uint64) bool
}

func (h *host) Create(ctx context.Context, host *asset.Host) {
	restfulx.ErrNotNilDebug(h.db.Create(host).Error, restfulx.OperatorErr)
}

func (h *host) Delete(ctx context.Context, hostId uint64) {
	restfulx.ErrNotNilDebug(h.db.Where("id = ?", hostId).Delete(&asset.Host{}).Error, restfulx.OperatorErr)
}

func (h *host) Update(ctx context.Context, hostId uint64, host *asset.Host) {
	restfulx.ErrNotNilDebug(h.db.Where("id = ?", hostId).Updates(host).Error, restfulx.OperatorErr)
}

func (h *host) List(ctx context.Context, page, limit int, groups []uint64, ip, label, description string) *form.PageHost {
	var (
		hostList []asset.Host
		total    int64
	)
	db := h.db

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

func (h *host) GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host) {
	host = &asset.Host{}
	restfulx.ErrNotNilDebug(h.db.Where("id = ?", hostId).Find(&host).Error, restfulx.OperatorErr)
	return host
}

func (h *host) GetByGroup(ctx context.Context, groups []uint64, page, limit int) *form.PageResult {
	var (
		hostList []*asset.Host
		total    int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(h.db.Where("group_id in (?)", groups).Find(&hostList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(h.db.Where("group_id in (?)", groups).Model(&asset.Host{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageResult{
			Data:  hostList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(h.db.Where("group_id in (?)", groups).Limit(limit).Offset((page-1)*limit).
		Find(&hostList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(h.db.Model(&asset.Host{}).Where("group_id in (?)", groups).Count(&total).Error, restfulx.OperatorErr)

	res := &form.PageResult{
		Data:  hostList,
		Total: total,
	}
	return res
}

func (h *host) UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(h.db.Where("id = ?", hostId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}

func (h *host) GetHostAccounts(ctx context.Context, hostId uint64) []*asset.Account {
	accountIds := make([]uint64, 0)
	list := h.rbac.List(ctx, 0, 0)
	for _, ha := range list.Data.([]*asset.HostAccount) {
		for _, ra := range ha.ResourceId {
			if ra == hostId {
				accountIds = append(accountIds, ha.AccountId...)
			}
		}
	}
	accountIds = RemoveRepByMap(accountIds)
	return h.account.GetByIds(ctx, accountIds)
}

func (h *host) CheckHostExist(ctx context.Context, hostId uint64) bool {
	host := &asset.Host{}
	err := h.db.Where("id = ?", hostId).Find(&host).Error
	if err == nil {
		return false
	}
	return true
}

func RemoveRepByMap(slc []uint64) []uint64 {
	var result []uint64
	tempMap := map[uint64]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
