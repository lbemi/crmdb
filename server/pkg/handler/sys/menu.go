package sys

import (
	"context"
	"github.com/lbemi/lbemi/pkg/bootstrap/log"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/model/sys"
	"github.com/lbemi/lbemi/pkg/services"
)

type MenuGetter interface {
	Menu() IMenu
}

// IMenu 菜单操作接口
type IMenu interface {
	Create(c context.Context, obj *form.MenusReq) (menu *sys.Menu, err error)
	Update(c context.Context, menu *form.UpdateMenusReq, menuID uint64) error
	Delete(c context.Context, menuID uint64) error
	Get(c context.Context, menuID uint64) (menu *sys.Menu, err error)
	List(c context.Context, page, limit int, menuType []int8) (res *form.PageMenu, err error)

	GetByIds(c context.Context, menuIDs []uint64) (menus *[]sys.Menu, err error)
	GetMenuByMenuNameUrl(context.Context, string, string) (*sys.Menu, error)
	CheckMenusIsExist(c context.Context, menuID uint64) bool
	UpdateStatus(c context.Context, menuID, status uint64) error
}

type menu struct {
	factory services.FactoryImp
}

func NewMenu(f services.FactoryImp) IMenu {
	return &menu{
		factory: f,
	}
}

func (m *menu) Create(c context.Context, obj *form.MenusReq) (menu *sys.Menu, err error) {
	if menu, err = m.factory.Menu().Create(&sys.Menu{
		Name:     obj.Name,
		Memo:     obj.Memo,
		ParentID: obj.ParentID,
		Status:   obj.Status,
		URL:      obj.URL,
		Icon:     obj.Icon,
		Sequence: obj.Sequence,
		MenuType: obj.MenuType,
		Method:   obj.Method,
		Code:     obj.Code,
	}); err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (m *menu) Update(c context.Context, menu *form.UpdateMenusReq, menuID uint64) error {
	res, err := m.factory.Menu().Get(menuID)
	if err != nil {
		log.Logger.Error(err)
		return err
	}
	err = m.factory.Menu().Update(menu, menuID)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	if res.URL != menu.URL || res.Method != menu.Method {
		err = m.factory.Authentication().UpdatePermissions(res.URL, res.Method, menu.URL, menu.Method)
		if err != nil {
			log.Logger.Error(err)
			return err
		}
	}
	return nil
}

func (m *menu) Delete(c context.Context, menuID uint64) error {
	menuInfo, err := m.factory.Menu().Get(menuID)
	// 如果报错或者未获取到menu信息则返回
	if err != nil || menuInfo == nil {
		log.Logger.Error(err)
		return err
	}

	// 清除rules
	err = m.factory.Authentication().DeleteRolePermission(menuInfo.URL, menuInfo.Method)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	// 清除menus
	err = m.factory.Menu().Delete(menuID)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	return nil
}

func (m *menu) Get(c context.Context, menuID uint64) (menu *sys.Menu, err error) {
	if menu, err = m.factory.Menu().Get(menuID); err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	return
}

func (m *menu) List(c context.Context, page, limit int, menuType []int8) (res *form.PageMenu, err error) {
	if res, err = m.factory.Menu().List(page, limit, menuType); err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (m *menu) GetByIds(c context.Context, menuIDs []uint64) (menus *[]sys.Menu, err error) {
	menus, err = m.factory.Menu().GetByIds(menuIDs)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	return
}

func (m *menu) GetMenuByMenuNameUrl(c context.Context, url, method string) (menu *sys.Menu, err error) {
	menu, err = m.factory.Menu().GetMenuByMenuNameUrl(url, method)
	return
}

func (m *menu) CheckMenusIsExist(c context.Context, menuID uint64) bool {
	_, err := m.factory.Menu().Get(menuID)
	if err != nil {
		log.Logger.Error(err)
		return false
	}
	return true
}

func (m *menu) UpdateStatus(c context.Context, menuID, status uint64) error {
	return m.factory.Menu().UpdateStatus(menuID, status)
}
