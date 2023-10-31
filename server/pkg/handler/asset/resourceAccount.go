package asset

import (
	"context"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
)

type ResourceAccountGetter interface {
	ResourceBindAccount() IResourceBindAccount
}

type IResourceBindAccount interface {
	BindAccount(ctx context.Context, hostAccount *asset.HostAccount)
	UnbindAccount(ctx context.Context, haId uint64)
	UpdateHostAccount(ctx context.Context, hostAccount *asset.HostAccount)
	List(ctx context.Context, page, limit int) *form.PageResult
	Get(ctx context.Context, haId uint64) *asset.HostAccount
}

type resourceBindAccount struct {
	factory services.Interface
}

func NewResourceBindAccount(f services.Interface) IResourceBindAccount {
	return &resourceBindAccount{
		factory: f,
	}
}

func (r *resourceBindAccount) BindAccount(ctx context.Context, hostAccount *asset.HostAccount) {
	r.factory.ResourceBindAccount().BindAccount(ctx, hostAccount)
}

func (r *resourceBindAccount) UnbindAccount(ctx context.Context, haId uint64) {
	r.factory.ResourceBindAccount().UnbindAccount(ctx, haId)
}

func (r *resourceBindAccount) UpdateHostAccount(ctx context.Context, hostAccount *asset.HostAccount) {
	r.factory.ResourceBindAccount().UpdateHostAccount(ctx, hostAccount)
}

func (r *resourceBindAccount) List(ctx context.Context, page, limit int) *form.PageResult {
	return r.factory.ResourceBindAccount().List(ctx, page, limit)
}

func (r *resourceBindAccount) Get(ctx context.Context, haId uint64) *asset.HostAccount {
	return r.factory.ResourceBindAccount().Get(ctx, haId)
}
