package asset

import (
	"context"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
)

type AccountGetter interface {
	Account() IAccount
}

type account struct {
	factory services.FactoryImp
}

func NewAccount(f services.FactoryImp) IAccount {
	return &account{
		factory: f,
	}
}

// IAccount 主机操作接口
type IAccount interface {
	Create(ctx context.Context, account *asset.Account)
	Delete(ctx context.Context, accountId uint64)
	Update(ctx context.Context, accountId uint64, account *asset.Account)
	List(ctx context.Context, page, limit int) *form.PageResult
	GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account)
	UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8)
	CheckAccountExist(ctx context.Context, accountId uint64) bool
}

func (m *account) Create(ctx context.Context, account *asset.Account) {
	m.factory.Account().Create(ctx, account)
}

func (m *account) Delete(ctx context.Context, accountId uint64) {
	m.factory.Account().Delete(ctx, accountId)

}

func (m *account) Update(ctx context.Context, accountId uint64, account *asset.Account) {
	m.factory.Account().Update(ctx, accountId, account)

}

func (m *account) List(ctx context.Context, page, limit int) *form.PageResult {
	return m.factory.Account().List(page, limit)

}

func (m *account) GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account) {
	return m.factory.Account().GetByAccountId(ctx, accountId)

}

func (m *account) UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8) {
	m.factory.Account().UpdateFiledStatus(ctx, accountId, updateFiled, status)

}

func (m *account) CheckAccountExist(ctx context.Context, accountId uint64) bool {
	h := m.factory.Account().GetByAccountId(ctx, accountId)
	if h == nil {
		return false
	}
	return true
}
