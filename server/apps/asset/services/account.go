package services

import (
	"context"
	"fmt"
	"github.com/lbemi/lbemi/apps/asset/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
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
	Create(ctx context.Context, account *entity.Account)
	Delete(ctx context.Context, accountId uint64)
	Update(ctx context.Context, accountId uint64, account *entity.Account)
	List(ctx context.Context, page, limit int, name, userName string) *entity2.PageResult
	GetByAccountId(ctx context.Context, accountId uint64) (account *entity.Account)
	UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8)
	CheckAccountExist(ctx context.Context, accountId uint64) bool
	GetByIds(ctx context.Context, accountIds []uint64) []*entity.Account
}

func (a *account) Create(ctx context.Context, account *entity.Account) {
	if account.Password != "" {
		account.Password = util.Encrypt(account.Password)
	}
	if account.Secret != "" {
		account.Password = util.Encrypt(account.Secret)
	}

	restfulx.ErrNotNilDebug(a.db.Create(account).Error, restfulx.OperatorErr)
}

func (a *account) Delete(ctx context.Context, accountId uint64) {
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Delete(&entity.Account{}).Error, restfulx.OperatorErr)
}

func (a *account) Update(ctx context.Context, accountId uint64, account *entity.Account) {
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

func (a *account) List(ctx context.Context, page, limit int, name, userName string) *entity2.PageResult {
	var (
		accountList []entity.Account
		total       int64
	)
	db := a.db
	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if userName != "" {
		db = db.Where("user_name LIKE ?", "%"+userName+"%")
	}

	restfulx.ErrNotNilDebug(db.Model(&entity.Account{}).Count(&total).Error, restfulx.OperatorErr)

	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(db.Find(&accountList).Error, restfulx.OperatorErr)
	} else {
		restfulx.ErrNotNilDebug(db.Limit(limit).Offset((page-1)*limit).Find(&accountList).Error, restfulx.OperatorErr)
	}

	res := &entity2.PageResult{
		Data:  accountList,
		Total: total,
	}
	return res
}

func (a *account) GetByAccountId(ctx context.Context, accountId uint64) (account *entity.Account) {
	account = &entity.Account{}
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Find(&account).Error, restfulx.OperatorErr)
	return account
}

func (a *account) UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(a.db.Where("id = ?", accountId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}

func (a *account) CheckAccountExist(ctx context.Context, accountId uint64) bool {
	account := &entity.Account{}
	err := a.db.Where("id = ?", accountId).Find(&account).Error
	if err == nil {
		return false
	}
	return true
}

func (a *account) GetByIds(ctx context.Context, accountIds []uint64) []*entity.Account {
	account := make([]*entity.Account, 0)
	restfulx.ErrNotNilDebug(a.db.Where("id in (?)", accountIds).Find(&account).Error, restfulx.OperatorErr)
	return account
}
