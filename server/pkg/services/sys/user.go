package sys

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"gorm.io/gorm"
)

type IUSer interface {
	Login(params *form.UserLoginForm) (user *sys.User, err error)
	Register(params *sys.User) (err error)
	Update(userID uint64, user *sys.User) (err error)
	GetUserInfoById(id uint64) (user *sys.User, err error)
	GetUserList(page, limit int) (*form.PageUser, error)
	DeleteUserByUserId(id uint64) error
	CheckUserExist(userName string) bool
	GetByName(name string) (*sys.User, error)
	GetRoleIdbyUser(userID uint64) (*[]sys.Role, error)
	SetUserRoles(userID uint64, roleIDs []uint64) error
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
	return
}

func (u *user) Register(params *sys.User) (err error) {
	return u.db.Create(&params).Error
}
func (u *user) Update(userID uint64, user *sys.User) (err error) {
	return u.db.Model(&sys.User{}).Where("id = ?", userID).Updates(&user).Error
}
func (u *user) GetUserInfoById(id uint64) (user *sys.User, err error) {
	err = u.db.First(&user, id).Error
	return
}

func (u *user) GetUserList(page, limit int) (*form.PageUser, error) {
	var (
		userList []sys.User
		total    int64
		err      error
	)

	// 全量查询
	if page == 0 && limit == 0 {
		if tx := u.db.Find(&userList); tx.Error != nil {
			return nil, tx.Error
		}

		if err := u.db.Model(&sys.User{}).Count(&total).Error; err != nil {
			return nil, err
		}

		res := &form.PageUser{
			Users: userList,
			Total: total,
		}
		return res, err
	}

	//分页数据
	if err := u.db.Limit(limit).Offset((page - 1) * limit).
		Find(&userList).Error; err != nil {
		return nil, err
	}

	if err := u.db.Model(&sys.User{}).Count(&total).Error; err != nil {
		return nil, err
	}

	res := &form.PageUser{
		Users: userList,
		Total: total,
	}
	return res, err

}

func (u *user) DeleteUserByUserId(userID uint64) error {
	return u.db.Where("id = ?", userID).Delete(&sys.User{}).Error
}

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
	if err := u.db.Where("name = ?", name).First(&obj).Error; err != nil {
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
		log.Logger.Errorf(err.Error())
		return nil, err
	}

	if roles != nil {
		res := GetTreeRoles(*roles, 0)
		return &res, err
	}

	return nil, err
}

// SetUserRoles 分配用户角色
func (u *user) SetUserRoles(userID uint64, roleIDS []uint64) (err error) {
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Errorf(err.(error).Error())
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Logger.Errorf(err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Where(&sys.UserRole{UserID: userID}).Delete(&sys.UserRole{}).Error; err != nil {
		log.Logger.Errorf(err.Error())
		tx.Rollback()
		return err
	}
	if len(roleIDS) > 0 {
		for _, rid := range roleIDS {
			rm := new(sys.UserRole)
			rm.RoleID = rid
			rm.UserID = userID
			if err := tx.Create(rm).Error; err != nil {
				log.Logger.Errorf(err.Error())
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// GetButtonsByUserID 获取菜单按钮
func (u *user) GetButtonsByUserID(userID uint64) (*[]sys.Menu, error) {
	var permissions []sys.Menu

	err := u.db.Debug().Table("menus").Select(" menus.id, menus.code,menus.menu_type,menus.status").
		Joins("left join role_menus on menus.id = role_menus.menu_id ").
		Joins("left join user_roles on user_roles.role_id = role_menus.role_id where role_menus.role_id in (?) and menus.menu_type in (2,3) and menus.status = 1",
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
	err := u.db.Debug().Table("menus").Select(" menus.id, menus.parent_id,menus.name,menus.memo, menus.path, menus.icon,menus.sequence,"+
		"menus.method, menus.menu_type, menus.status, menus.component, menus.title, menus.isLink,menus.isHide,menus.isAffix,menus.isKeepAlive,menus.isIframe").
		Joins("left join role_menus on menus.id = role_menus.menu_id where role_menus.role_id in (?) and menus.menu_type = 1 and menus.status = 1",
			u.db.Table("roles").Select("roles.id").
				Joins("left join user_roles on user_roles.role_id = roles.id where  user_roles.user_id = ? and roles.status = 1", userID)).
		Group("id").
		Order("parent_id ASC").
		Order("sequence DESC").
		Scan(&menus).Error

	if err != nil {
		log.Logger.Error(err)
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
