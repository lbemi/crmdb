package services

import (
	"context"
	entity2 "github.com/lbemi/lbemi/apps/log/entity"
	operatorLog "github.com/lbemi/lbemi/apps/log/services"
	"github.com/lbemi/lbemi/apps/system/api/form"
	"github.com/lbemi/lbemi/apps/system/entity"
	"github.com/lbemi/lbemi/pkg/bootstrap/policy"
	entity3 "github.com/lbemi/lbemi/pkg/common/entity"
	"gorm.io/gorm"
	"time"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	"github.com/mssola/useragent"
)

type UserGetter interface {
	User() IUSer
}

type IUSer interface {
	Login(rc *rctx.ReqCtx, params *form.UserLoginForm) (user *entity.User)
	Register(c context.Context, params *form.RegisterUserForm)
	Update(c context.Context, userID uint64, params *form.UpdateUserFrom)
	GetUserInfoById(c context.Context, id uint64) (user *entity.User)
	GetUserList(c context.Context, param *entity3.PageParam, condition *entity.User) *form.PageUser
	DeleteUserByUserId(c context.Context, id uint64)
	CheckUserExist(c context.Context, userName string) bool
	GetByName(c context.Context, name string) *entity.User
	GetRoleIDByUser(c context.Context, userID uint64) []*entity.Role
	SetUserRoles(c context.Context, userID uint64, roleIDs []uint64)
	GetButtonsByUserID(c context.Context, userID uint64) *[]string
	GetLeftMenusByUserID(c context.Context, userID uint64) *[]entity.Menu
	UpdateStatus(c context.Context, userID, status uint64)
}

type User struct {
	db     *gorm.DB
	policy policy.IPolicy
	menu   IMenu
	log    operatorLog.ILoginLog
}

func NewUser(db *gorm.DB, policy policy.IPolicy, menu IMenu, log operatorLog.ILoginLog) IUSer {
	return &User{
		db:     db,
		policy: policy,
		menu:   menu,
		log:    log,
	}
}

func (u *User) Login(rc *rctx.ReqCtx, params *form.UserLoginForm) (user *entity.User) {
	err := u.db.Where("user_name = ?", params.UserName).First(&user).Error
	restfulx.ErrNotNilDebug(err, restfulx.PasswdWrong)

	pass := util.BcryptMakeCheck([]byte(params.Password), user.Password)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				switch t := r.(type) {
				case *restfulx.OpsError:
					log.Logger.Error(t.Error())
				case error:
					log.Logger.Error(t)
				case string:
					log.Logger.Error(t)
				}
			}
		}()
		req := rc.Request.Request
		ua := useragent.New(req.UserAgent())
		bName, bVersion := ua.Browser()
		log := &entity2.LogLogin{
			Username:      params.UserName,
			Ipaddr:        rc.ClientIP(),
			LoginLocation: "",
			Browser:       bName + ":" + bVersion,
			Os:            ua.OS(),
			Platform:      ua.Platform(),
			LoginTime:     time.Now(),
			Remark:        req.UserAgent(),
		}
		log.LoginLocation = util.GetRealAddressByIP(log.Ipaddr)
		if pass && err == nil {
			log.Status = "1"
			log.Msg = "登录成功"
		} else {
			log.Status = "-1"
			log.Msg = "登录失败"
		}
		u.log.Add(log)
	}()
	restfulx.ErrNotTrue(user.Status == 1, restfulx.UserDeny)
	restfulx.ErrNotTrue(pass, restfulx.PasswdWrong)
	return user
}

func (u *User) Register(ctx context.Context, params *form.RegisterUserForm) {
	userInfo := &entity.User{
		UserName: params.UserName,
		Password: util.BcryptMake([]byte(params.Password)),
		//Mobile:      params.Mobile,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}

	restfulx.ErrNotTrue(!u.CheckUserExist(ctx, params.UserName), restfulx.UserExist)
	restfulx.ErrNotNilDebug(u.db.Create(&userInfo).Error, restfulx.OperatorErr)
}

func (u *User) Update(ctx context.Context, userID uint64, params *form.UpdateUserFrom) {
	userInfo := &entity.User{
		UserName:    params.UserName,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}
	restfulx.ErrNotNilDebug(u.db.Model(&entity.User{}).Where("id = ?", userID).Updates(userInfo).Error, restfulx.OperatorErr)
}

func (u *User) GetUserInfoById(ctx context.Context, id uint64) *entity.User {
	var user *entity.User
	err := u.db.First(&user, id).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return user
}

func (u *User) GetUserList(ctx context.Context, pageParam *entity3.PageParam, condition *entity.User) *form.PageUser {
	db := u.db
	if condition.Status != 0 {
		db = db.Where("status = ?", condition.Status)
	}

	if condition.UserName != "" {
		db = db.Where("user_name like ?", "%"+condition.UserName+"%")
	}
	var (
		userList []entity.User
		total    int64
	)

	// 全量查询
	if pageParam.Page == 0 && pageParam.Limit == 0 {
		err := db.Find(&userList).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		err = db.Model(&entity.User{}).Count(&total).Error
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		res := &form.PageUser{
			Users: userList,
			Total: total,
		}
		return res
	}
	//分页数据
	err := db.Limit(pageParam.Limit).Offset((pageParam.Page - 1) * pageParam.Limit).
		Find(&userList).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	err = db.Model(&entity.User{}).Count(&total).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	res := &form.PageUser{
		Users: userList,
		Total: total,
	}
	return res
}

func (u *User) DeleteUserByUserId(ctx context.Context, id uint64) {
	//TODO 修复一下逻辑
	tx := u.db.Where("id = ?", id).Delete(&entity.User{})
	restfulx.ErrNotNilDebug(tx.Error, restfulx.OperatorErr)
	u.policy.DeleteUser(id)
}

func (u *User) CheckUserExist(ctx context.Context, userName string) bool {
	err := u.db.Where("user_name = ?", userName).First(&entity.User{}).Error
	if err != nil {
		return false
	}
	return true
}

func (u *User) GetByName(ctx context.Context, name string) *entity.User {
	var obj *entity.User
	err := u.db.Where("name = ?", name).First(&obj).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return obj
}

// GetRoleIDByUser 查询用户角色
func (u *User) GetRoleIDByUser(ctx context.Context, userID uint64) []*entity.Role {
	var roles []*entity.Role
	subRoleIdSql := u.db.Select("role_id").Where("user_id = ?", userID).Table("user_roles")
	err := u.db.Table("roles").
		Select("roles.*").
		Joins("left join user_roles on roles.id = user_roles.role_id").
		Where("roles.id in (?)", subRoleIdSql).
		Or("roles.parent_id in (?)", subRoleIdSql).
		Group("id").
		Order("id asc").
		Order("sequence desc").
		Scan(&roles).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	if roles != nil {
		res := GetTreeRoles(roles, 0)
		return res
	}

	return roles
}

// SetUserRoles 分配用户角色
func (u *User) SetUserRoles(ctx context.Context, userID uint64, roleIDS []uint64) {

	// 配置role_users表
	tx := u.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	if err := tx.Where(&entity.UserRole{UserID: userID}).Delete(&entity.UserRole{}).Error; err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}
	if len(roleIDS) > 0 {
		for _, rid := range roleIDS {
			rm := new(entity.UserRole)
			rm.RoleID = rid
			rm.UserID = userID
			if err := tx.Create(rm).Error; err != nil {
				tx.Rollback()
				restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
			}
		}
	}
	err := tx.Commit().Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	// 添加规则到rules表
	err = u.policy.AddRoleForUser(ctx, userID, roleIDS)
	if err != nil {
		tx.Rollback()
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

}

// GetButtonsByUserID 获取菜单按钮
func (u *User) GetButtonsByUserID(ctx context.Context, userID uint64) *[]string {
	var menus []*entity.Menu
	err := u.db.Table("menus").Select(" menus.id, menus.code,menus.menuType,menus.status").
		Joins("left join role_menus on menus.id = role_menus.menuID ").
		Joins("left join user_roles on user_roles.role_id = role_menus.roleID where role_menus.roleID in (?) and menus.menuType in (2,3) and menus.status = 1",
			u.db.Table("roles").Select("roles.id").
				Joins("left join user_roles on user_roles.role_id = roles.id where  user_roles.user_id = ? and roles.status = 1", userID)).
		Group("id").
		Scan(&menus).Error

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	var res []string
	for _, v := range menus {
		if v.Code != "" {
			res = append(res, v.Code)
		}
	}
	return &res
}

// GetLeftMenusByUserID 根据用户ID获取左侧菜单
func (u *User) GetLeftMenusByUserID(ctx context.Context, userID uint64) *[]entity.Menu {
	var menus []entity.Menu
	err := u.db.Table("menus").Select(" menus.id, menus.parentID,menus.name,menus.memo, menus.path, menus.icon,menus.sequence,"+
		"menus.method, menus.menuType, menus.status,menus.redirect, menus.component, menus.isK8s,menus.title, menus.isLink,menus.isHide,menus.isAffix,menus.isKeepAlive,menus.isIframe").
		Joins("left join role_menus on menus.id = role_menus.menuID where role_menus.roleID in (?) and menus.menuType = 1 and menus.status = 1",
			u.db.Table("roles").Select("roles.id").
				Joins("left join user_roles on user_roles.role_id = roles.id where  user_roles.user_id = ? and roles.status = 1", userID)).
		Group("id").
		Order("parentID ASC").
		Order("sequence DESC").
		Scan(&menus).Error

	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	if len(menus) == 0 {
		return &menus
	}
	treeMenusList := GetTreeMenus(menus, 0)
	return &treeMenusList
}

func (u *User) UpdateStatus(ctx context.Context, userId, status uint64) {
	restfulx.ErrNotNilDebug(u.db.Model(&entity.User{}).Where("id = ?", userId).Update("status", status).Error, restfulx.OperatorErr)
}
