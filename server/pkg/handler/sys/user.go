package sys

import (
	operatorLog "github.com/lbemi/lbemi/pkg/handler/logsys"
	"github.com/lbemi/lbemi/pkg/handler/policy"
	"gorm.io/gorm"
	"time"

	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/logsys"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/rctx"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/util"
	"github.com/mssola/useragent"
)

type UserGetter interface {
	User() IUSer
}

type IUSer interface {
	Login(rc *rctx.ReqCtx, params *form.UserLoginForm) (user *sys.User)
	Register(params *form.RegisterUserForm)
	Update(userID uint64, params *form.UpdateUserFrom)
	GetUserInfoById(id uint64) (user *sys.User)
	GetUserList(param *model.PageParam, condition *sys.User) *form.PageUser
	DeleteUserByUserId(id uint64)
	CheckUserExist(userName string) bool
	GetByName(name string) *sys.User
	GetRoleIDByUser(userID uint64) *[]sys.Role
	SetUserRoles(userID uint64, roleIDs []uint64)
	GetButtonsByUserID(userID uint64) *[]string
	GetLeftMenusByUserID(userID uint64) *[]sys.Menu
	UpdateStatus(userID, status uint64)
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

func (u *User) Login(rc *rctx.ReqCtx, params *form.UserLoginForm) (user *sys.User) {
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
		log := &logsys.LogLogin{
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

func (u *User) Register(params *form.RegisterUserForm) {
	userInfo := &sys.User{
		UserName: params.UserName,
		Password: util.BcryptMake([]byte(params.Password)),
		//Mobile:      params.Mobile,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}

	restfulx.ErrNotTrue(!u.CheckUserExist(params.UserName), restfulx.UserExist)
	restfulx.ErrNotNilDebug(u.db.Create(&userInfo).Error, restfulx.OperatorErr)
}

func (u *User) Update(userID uint64, params *form.UpdateUserFrom) {
	userInfo := &sys.User{
		UserName:    params.UserName,
		Email:       params.Email,
		Description: params.Description,
		Status:      params.Status,
	}
	restfulx.ErrNotNilDebug(u.db.Model(&sys.User{}).Where("id = ?", userID).Updates(userInfo).Error, restfulx.OperatorErr)
}

func (u *User) GetUserInfoById(id uint64) *sys.User {
	var user *sys.User
	err := u.db.First(&user, id).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return user
}

func (u *User) GetUserList(pageParam *model.PageParam, condition *sys.User) *form.PageUser {
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
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		err = db.Model(&sys.User{}).Count(&total).Error
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

	err = db.Model(&sys.User{}).Count(&total).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	res := &form.PageUser{
		Users: userList,
		Total: total,
	}
	return res
}

func (u *User) DeleteUserByUserId(id uint64) {
	//TODO 修复一下逻辑
	tx := u.db.Where("id = ?", id).Delete(&sys.User{})
	restfulx.ErrNotNilDebug(tx.Error, restfulx.OperatorErr)
	u.policy.DeleteUser(id)
}

func (u *User) CheckUserExist(userName string) bool {
	err := u.db.Where("user_name = ?", userName).First(&sys.User{}).Error
	if err != nil {
		return false
	}
	return true
}

func (u *User) GetByName(name string) *sys.User {
	var obj *sys.User
	err := u.db.Where("name = ?", name).First(&obj).Error
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return obj
}

// GetRoleIDByUser 查询用户角色
func (u *User) GetRoleIDByUser(userID uint64) *[]sys.Role {
	var roles *[]sys.Role
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
		res := GetTreeRoles(*roles, 0)
		return &res
	}

	return roles
}

// SetUserRoles 分配用户角色
func (u *User) SetUserRoles(userID uint64, roleIDS []uint64) {

	// 配置role_users表
	tx, err := u.factory.User().SetUserRoles(userID, roleIDS)
	if err != nil {
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

	// 添加规则到rules表
	err = u.factory.Authentication().AddRoleForUser(userID, roleIDS)
	if err != nil {
		// 报错则回退数据
		tx.Rollback()
		for _, roleId := range roleIDS {
			u.factory.Authentication().DeleteRoleWithUser(userID, roleId)
		}
		restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	}

}

// GetButtonsByUserID 获取菜单按钮
func (u *User) GetButtonsByUserID(userID uint64) *[]string {
	menus, err := u.factory.User().GetButtonsByUserID(userID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	var res []string
	for _, v := range *menus {
		if v.Code != "" {
			res = append(res, v.Code)
		}
	}

	return &res
}

// GetLeftMenusByUserID 根据用户ID获取左侧菜单
func (u *User) GetLeftMenusByUserID(userID uint64) *[]sys.Menu {
	res, err := u.factory.User().GetLeftMenusByUserID(userID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (u *User) UpdateStatus(userId, status uint64) {
	restfulx.ErrNotNilDebug(u.factory.User().UpdateStatus(userId, status), restfulx.OperatorErr)
}
