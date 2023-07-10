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
	List(ctx context.Context, page, limit int, name, userName string) *form.PageResult
	GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account)
	GetByIds(ctx context.Context, ids []uint64) []*asset.Account
	UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8)
}

// Create creates a new account.
//
// ctx - the context for the operation.
// account - the account to be created.
func (m *account) Create(ctx context.Context, account *asset.Account) {
	restfulx.ErrNotNilDebug(m.db.Create(account).Error, restfulx.OperatorErr)
}

// Delete deletes the account with the specified accountId.
//
// Parameters:
// - ctx: The context.Context object.
// - accountId: The ID of the account to be deleted.
func (m *account) Delete(ctx context.Context, accountId uint64) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Delete(&asset.Account{}).Error, restfulx.OperatorErr)
}

func (m *account) Update(ctx context.Context, accountId uint64, account *asset.Account) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Updates(account).Error, restfulx.OperatorErr)
}

func (m *account) List(ctx context.Context, page, limit int, name, userName string) *form.PageResult {
	var (
		accountList []asset.Account
		total       int64
	)

	db := m.db

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

// GetByAccountId retrieves an account by its account ID.
//
// ctx: The context.Context object for the request.
// accountId: The ID of the account to retrieve.
// Returns:
//   - account: The account object with the specified ID.
func (m *account) GetByAccountId(ctx context.Context, accountId uint64) (account *asset.Account) {
	account = &asset.Account{}
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Find(&account).Error, restfulx.OperatorErr)
	return account
}

// GetByIds returns an array of account objects based on the provided account IDs.
//
// The function takes two parameters:
// - ctx: the context.Context object for handling cancellation and timeouts.
// - accountIds: a slice of uint64 values representing the account IDs.
//
// The function returns an array of *asset.Account objects.
func (m *account) GetByIds(ctx context.Context, accountIds []uint64) []*asset.Account {
	account := make([]*asset.Account, 0)
	restfulx.ErrNotNilDebug(m.db.Where("id in (?)", accountIds).Find(&account).Error, restfulx.OperatorErr)
	return account
}

// UpdateFiledStatus updates the filed status of an account.
//
// It takes the following parameters:
// - ctx: the context.Context object for the operation.
// - accountId: the ID of the account.
// - updateFiled: the field to be updated.
// - status: the new status value.
func (m *account) UpdateFiledStatus(ctx context.Context, accountId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(m.db.Where("id = ?", accountId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}
