package asset

import (
	"context"
	"fmt"
	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
)

type AccountGetter interface {
	Account() IAccount
}

type account struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) IAccount {
	return &account{
		db: db,
	}
}

// IAccount 主机操作接口
type IAccount interface {
	Create(ctx context.Context, account *asset.Account)
	Delete(ctx context.Context, accountId uint64)
	Update(ctx context.Context, accountId uint64, account *asset.Account)
	List(ctx context.Context, page, limit int, name, userName string) *form.PageResult
	GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account)
	UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8)
	CheckAccountExist(ctx context.Context, accountId uint64) bool
	GetByIds(ctx context.Context, accountIds []uint64) []*asset.Account
}

func (a *account) Create(ctx context.Context, account *asset.Account) {
	if account.Password != "" {
		account.Password = util.Encrypt(account.Password)
	}
	if account.Secret != "" {
		account.Password = util.Encrypt(account.Secret)
	}

	restfulx.ErrNotNilDebug(a.db.Create(account).Error, restfulx.OperatorErr)
}

func (a *account) Delete(ctx context.Context, accountId uint64) {
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Delete(&asset.Account{}).Error, restfulx.OperatorErr)
}

func (a *account) Update(ctx context.Context, accountId uint64, account *asset.Account) {
	exist := a.CheckAccountExist(ctx, accountId)
	if !exist {
		restfulx.ErrNotNil(fmt.Errorf("账户不存在"), restfulx.OperatorErr)
	}
	if account.Password != "" {
		account.Password = util.Encrypt(account.Password)
	}
	if account.Secret != "" {
		account.Password = util.Encrypt(account.Secret)
	}
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Updates(account).Error, restfulx.OperatorErr)
}

func (a *account) List(ctx context.Context, page, limit int, name, userName string) *form.PageResult {
	var (
		accountList []asset.Account
		total       int64
	)
	db := a.db
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if userName != "" {
		db = db.Where("user_name LIKE ?", "%"+userName+"%")
	}

	restfulx.ErrNotNilDebug(db.Model(&asset.Account{}).Count(&total).Error, restfulx.OperatorErr)

	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(db.Find(&accountList).Error, restfulx.OperatorErr)
	} else {
		restfulx.ErrNotNilDebug(db.Limit(limit).Offset((page-1)*limit).Find(&accountList).Error, restfulx.OperatorErr)
	}

	res := &form.PageResult{
		Data:  accountList,
		Total: total,
	}
	return res
}

func (a *account) GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account) {
	account = &asset.Account{}
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Find(&account).Error, restfulx.OperatorErr)
	return account
}

func (a *account) UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}

func (a *account) CheckAccountExist(ctx context.Context, accountId uint64) bool {
	account := &asset.Account{}
	err := a.db.Where("id = ?", accountId).Find(&account).Error
	if err == nil {
		return false
	}
	return true
}

func (a *account) GetByIds(ctx context.Context, accountIds []uint64) []*asset.Account {
	account := make([]*asset.Account, 0)
	restfulx.ErrNotNilDebug(a.db.Where("id in (?)", accountIds).Find(&account).Error, restfulx.OperatorErr)
	return account
}
