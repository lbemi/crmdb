package sys

import (
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"github.com/lbemi/lbemi/pkg/services"
)

type MenuGetter interface {
	Menu() IMenu
}

// IMenu 菜单操作接口
type IMenu interface {
	Create(obj *form.MenusReq) *sys.Menu
	Update(menu *form.UpdateMenusReq, menuID uint64)
	Delete(menuID uint64)
	Get(menuID uint64) *sys.Menu
	List(page, limit int, menuType []int8, isTree bool) *form.PageMenu

	GetByIds(menuIDs []uint64) *[]sys.Menu
	GetMenuByMenuNameUrl(string, string) *sys.Menu
	CheckMenusIsExist(menuID uint64) bool
	UpdateStatus(menuID, status uint64)
}

type menu struct {
	factory services.FactoryImp
}

func NewMenu(f services.FactoryImp) IMenu {
	return &menu{
		factory: f,
	}
}

func (m *menu) Create(obj *form.MenusReq) *sys.Menu {
	_, err := m.factory.Menu().GetMenuByMenuNameUrl(obj.Path, obj.Method)
	if err.Error() != "record not found" {
		restfulx.ErrNotNilDebug(err, restfulx.ResourceExist)
	}
	res, err := m.factory.Menu().Create(&sys.Menu{
		Name:     obj.Name,
		Memo:     obj.Memo,
		ParentID: obj.ParentID,
		Status:   obj.Status,
		Path:     obj.Path,
		Group:    obj.Group,
		Meta: sys.Meta{
			Icon:        obj.Meta.Icon,
			Title:       obj.Meta.Title,
			IsLink:      obj.Meta.IsLink,
			IsAffix:     obj.Meta.IsAffix,
			IsHide:      obj.Meta.IsHide,
			IsIframe:    obj.Meta.IsIframe,
			IsKeepAlive: obj.Meta.IsKeepAlive,
			IsK8s:       obj.Meta.IsK8s,
		},
		Redirect:  obj.Redirect,
		Component: obj.Component,
		Sequence:  obj.Sequence,
		MenuType:  obj.MenuType,
		Method:    obj.Method,
		Code:      obj.Code,
	})
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (m *menu) Update(menu *form.UpdateMenusReq, menuID uint64) {
	restfulx.ErrNotTrue(m.CheckMenusIsExist(menuID), restfulx.ResourceExist)

	res, err := m.factory.Menu().Get(menuID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	err = m.factory.Menu().Update(menu, menuID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)

	if res.Path != menu.Path || res.Method != menu.Method {
		err = m.factory.Authentication().UpdatePermissions(res.Path, res.Method, menu.Path, menu.Method)
		if err != nil {
			restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
		}
	}
}

func (m *menu) Delete(menuID uint64) {
	restfulx.ErrNotTrue(m.CheckMenusIsExist(menuID), restfulx.ResourceExist)
	menuInfo, err := m.factory.Menu().Get(menuID)
	// 如果报错或者未获取到menu信息则返回
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	// 清除rules
	err = m.factory.Authentication().DeleteRolePermission(menuInfo.Path, menuInfo.Method)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	// 清除menus
	err = m.factory.Menu().Delete(menuID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
}

func (m *menu) Get(menuID uint64) *sys.Menu {
	res, err := m.factory.Menu().Get(menuID)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (m *menu) List(page, limit int, menuType []int8, isTree bool) *form.PageMenu {
	res, err := m.factory.Menu().List(page, limit, menuType, isTree)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (m *menu) GetByIds(menuIDs []uint64) *[]sys.Menu {
	res, err := m.factory.Menu().GetByIds(menuIDs)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (m *menu) GetMenuByMenuNameUrl(url, method string) *sys.Menu {
	res, err := m.factory.Menu().GetMenuByMenuNameUrl(url, method)
	restfulx.ErrNotNilDebug(err, restfulx.OperatorErr)
	return res
}

func (m *menu) CheckMenusIsExist(menuID uint64) bool {
	_, err := m.factory.Menu().Get(menuID)
	if err != nil {
		log.Logger.Error(err)
		return false
	}
	return true
}

func (m *menu) UpdateStatus(menuID, status uint64) {
	restfulx.ErrNotNilDebug(m.factory.Menu().UpdateStatus(menuID, status), restfulx.OperatorErr)
}
