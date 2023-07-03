package asset

import (
	"context"
	"github.com/lbemi/lbemi/pkg/restfulx"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"gorm.io/gorm"
)

type account struct {
	db *gorm.DB
}

func NewAccount(DB *gorm.DB) IAccount {
	return &account{db: DB}
}

type IAccount interface {
	Create(ctx context.Context, account *asset.Account)
	Delete(ctx context.Context, accountId uint64)
	Update(ctx context.Context, accountId uint64, account *asset.Account)
	List(page, limit int) *form.PageResult
	GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account)
	UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8)
}

func (m *account) Create(ctx context.Context, account *asset.Account) {
	restfulx.ErrNotNilDebug(m.db.Create(account).Error, restfulx.OperatorErr)
}

func (m *account) Delete(ctx context.Context, accountId uint64) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Delete(&asset.Account{}).Error, restfulx.OperatorErr)
}

func (m *account) Update(ctx context.Context, accountId uint64, account *asset.Account) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Updates(account).Error, restfulx.OperatorErr)
}

func (m *account) List(page, limit int) *form.PageResult {
	var (
		accountList []asset.Account
		total       int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(m.db.Find(&accountList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(m.db.Model(&asset.Account{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageResult{
			Data:  accountList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(m.db.Limit(limit).Offset((page-1)*limit).
		Find(&accountList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(m.db.Model(&asset.Account{}).Count(&total).Error, restfulx.OperatorErr)

	res := &form.PageResult{
		Data:  accountList,
		Total: total,
	}
	return res
}

func (m *account) GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account) {
	account = &asset.Account{}
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Find(&account).Error, restfulx.OperatorErr)
	return account
}

func (m *account) UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}
