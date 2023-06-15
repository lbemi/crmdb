package sys

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"

	"gorm.io/gorm"
)

type IUSer interface {
	Login(params *form.UserLoginForm) (user *sys.User, err error)
	Register(params *sys.User) error
	Update(userID uint64, user *sys.User) error
	GetUserInfoById(id uint64) (user *sys.User, err error)
	GetUserList(pageParam *model.PageParam, condition *sys.User) (*form.PageUser, error)
	DeleteUserByUserId(id uint64) error
	CheckUserExist(userName string) bool
	GetByName(name string) (*sys.User, error)
	GetRoleIdbyUser(userID uint64) (*[]sys.Role, error)
	SetUserRoles(userID uint64, roleIDs []uint64) (*gorm.DB, error)
	GetButtonsByUserID(userID uint64) (*[]sys.Menu, error)
	GetLeftMenusByUserID(userID uint64) (*[]sys.Menu, error)
	UpdateStatus(userID, status uint64) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUSer {
	return &user{db: db}
}

func (u *user) Login(params *form.UserLoginForm) (user *sys.User, err error) {
	err = u.db.Where("user_name = ?", params.UserName).First(&user).Error
	return user, err
}

func (u *user) Register(params *sys.User) error {
	return u.db.Create(&params).Error
}

func (u *user) Update(userID uint64, user *sys.User) error {
	return u.db.Model(&sys.User{}).Where("id = ?", userID).Updates(&user).Error
}
func (u *user) GetUserInfoById(id uint64) (user *sys.User, err error) {
	err = u.db.First(&user, id).Error
	return
}

func (u *user) GetUserList(pageParam *model.PageParam, condition *sys.User) (*form.PageUser, error) {
	db := u.db
	if condition.Status != 0 {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.UserName != "" {
		db = db.Where("user_name like ?", "%"+condition.UserName+"%")
	}
	var (
		userList []sys.User
		total    int64
	)

	// 全量查询
	if pageParam.Page == 0 && pageParam.Limit == 0 {
		err := db.Find(&userList).Error
		if err != nil {
			return nil, err
		}
		err = db.Model(&sys.User{}).Count(&total).Error

		if err != nil {
			return nil, err
		}

		res := &form.PageUser{
			Users: userList,
			Total: total,
		}
		return res, nil
	}

	//分页数据
	err := db.Limit(pageParam.Limit).Offset((pageParam.Page - 1) * pageParam.Limit).
		Find(&userList).Error

	if err != nil {
		return nil, err
	}

	err = db.Model(&sys.User{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	res := &form.PageUser{
		Users: userList,
		Total: total,
	}

	return res, nil
}

func (u *user) DeleteUserByUserId(userID uint64) error {
	return u.db.Where("id = ?", userID).Delete(&sys.User{}).Error
}

// CheckUserExist 检查用户是否存在，存在返回true，否则false
func (u *user) CheckUserExist(userName string) bool {
	err := u.db.Where("user_name = ?", userName).First(&sys.User{}).Error
	if err != nil {
		return false
	}

	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (u *user) GetByName(name string) (*sys.User, error) {
	var obj sys.User
	err := u.db.Where("name = ?", name).First(&obj).Error
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// GetRoleIdbyUser 查询用户角色
func (u *user) GetRoleIdbyUser(userID uint64) (roles *[]sys.Role, err error) {
	subRoleIdSql := u.db.Select("role_id").Where("user_id = ?", userID).Table("user_roles")
	err = u.db.Table("roles").
		Select("roles.*").
		Joins("left join user_roles on roles.id = user_roles.role_id").
		Where("roles.id in (?)", subRoleIdSql).
		Or("roles.parent_id in (?)", subRoleIdSql).
		Group("id").
		Order("id asc").
		Order("sequence desc").
		Scan(&roles).Error
	if err != nil {
		return nil, err
	}

	if roles != nil {
		res := GetTreeRoles(*roles, 0)
		return &res, nil
	}

	return
}

// SetUserRoles 分配用户角色
func (u *user) SetUserRoles(userID uint64, roleIDS []uint64) (*gorm.DB, error) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Logger.Errorf(err.Error())
		tx.Rollback()
		return tx, err
	}

	if err := tx.Where(&sys.UserRole{UserID: userID}).Delete(&sys.UserRole{}).Error; err != nil {
		log.Logger.Errorf(err.Error())
		tx.Rollback()
		return tx, err
	}
	if len(roleIDS) > 0 {
		for _, rid := range roleIDS {
			rm := new(sys.UserRole)
			rm.RoleID = rid
			rm.UserID = userID
			if err := tx.Create(rm).Error; err != nil {
				log.Logger.Errorf(err.Error())
				tx.Rollback()
				return tx, err
			}
		}
	}
	err := tx.Commit().Error
	if err != nil {
		return tx, err
	}
	return tx, nil
}

// GetButtonsByUserID 获取菜单按钮
func (u *user) GetButtonsByUserID(userID uint64) (*[]sys.Menu, error) {
	var permissions []sys.Menu

	err := u.db.Table("menus").Select(" menus.id, menus.code,menus.menuType,menus.status").
		Joins("left join role_menus on menus.id = role_menus.menuID ").
		Joins("left join user_roles on user_roles.role_id = role_menus.roleID where role_menus.roleID in (?) and menus.menuType in (2,3) and menus.status = 1",
			u.db.Table("roles").Select("roles.id").
				Joins("left join user_roles on user_roles.role_id = roles.id where  user_roles.user_id = ? and roles.status = 1", userID)).
		Group("id").
		Scan(&permissions).Error

	if err != nil {
		return nil, err
	}
	return &permissions, nil
}

// GetLeftMenusByUserID 根据用户ID获取左侧菜单
func (u *user) GetLeftMenusByUserID(userID uint64) (*[]sys.Menu, error) {
	var menus []sys.Menu
	err := u.db.Table("menus").Select(" menus.id, menus.parentID,menus.name,menus.memo, menus.path, menus.icon,menus.sequence,"+
		"menus.method, menus.menuType, menus.status,menus.redirect, menus.component, menus.isK8s,menus.title, menus.isLink,menus.isHide,menus.isAffix,menus.isKeepAlive,menus.isIframe").
		Joins("left join role_menus on menus.id = role_menus.menuID where role_menus.roleID in (?) and menus.menuType = 1 and menus.status = 1",
			u.db.Table("roles").Select("roles.id").
				Joins("left join user_roles on user_roles.role_id = roles.id where  user_roles.user_id = ? and roles.status = 1", userID)).
		Group("id").
		Order("parentID ASC").
		Order("sequence DESC").
		Scan(&menus).Error

	if err != nil {
		return nil, err
	}

	if len(menus) == 0 {
		return &menus, nil
	}
	treeMenusList := GetTreeMenus(menus, 0)
	return &treeMenusList, nil
}

func (u *user) UpdateStatus(userId, status uint64) error {
	return u.db.Model(&sys.User{}).Where("id = ?", userId).Update("status", status).Error
}
