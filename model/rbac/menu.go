package rbac

import (
	"encoding/json"
	"github.com/ahmetb/go-linq"
	"github.com/xilloader/camp/model"
	"io/ioutil"
	"log"
	"path"
	"reflect"
)

type Menu struct {
	model.IoModel
	MenuId   string `json:"menu_id" gorm:"primary_key;type:varchar(128);default:'';comment:'菜单id'"`
	ParentID string `json:"parent_id" gorm:"type:varchar(128);default:'';comment:'父级菜单id'"`
	Name     string `json:"name" gorm:"type:varchar(128);index:idx_name;default:'';comment:'菜单名称'" json:"name"`
	Path     string `json:"path" gorm:"type:varchar(128);index:idx_method_path;default:'';comment:'菜单路径'"`
	Method   string `json:"method" gorm:"type:varchar(128);index:idx_method_path;default:'';comment:'请求方法'"`
	File     string `json:"file" gorm:"type:varchar(128);default:'';comment:'页面文件路径'"`
	Code     string `json:"code" gorm:"type:varchar(128);default:'';comment:'菜单代码';unique_index:uk_menu_code;"`
	Status   uint8  `json:"status" gorm:"type:tinyint(1);default:1;comment:'状态 (1:启用 2:不启用)'"`
	Sequence uint8  `json:"sequence" gorm:"type:int;default:0;comment:'排序值'" json:"sequence" form:"sequence"` // 同级别排序
	Field    uint32 `json:"field" gorm:"type:int;default:'';comment:'菜单所属领域'"`                                // 不同的客户端对应不同菜单 按二进制位区分
	MenuType uint8  `json:"menu_type" gorm:"column:menu_type;type:tinyint(1);" json:"menu_type" form:"menu_type"`
	Icon     string `json:"icon" gorm:"type:varchar(128);default:'';comment:'图标'"`
	Memo     string `json:"memo" gorm:"type:varchar(256);default:'';comment:'备注 '"`
	Children []Menu `json:"children" gorm:"-"`
	Chosen   bool   `json:"chosen" gorm:"-"`
}

const MenuTypeModule = 1
const MenuTypeMenu = 2
const MenuTypeOperate = 3

var MenuTypeMap = map[uint8]bool{
	MenuTypeModule:  true,
	MenuTypeMenu:    true,
	MenuTypeOperate: true,
}

type MenuMeta struct {
	Title string `json:"title"` // 标题
	Icon  string `json:"icon"`  // 图标
	File  string `json:"file"`  // 文件路径
	Cache bool   `json:"-"`     // 是不是缓存
}

type MenuModel struct {
	MenuId    string      `json:"menu_id"`
	Path      string      `json:"path"`                // 路由
	Name      string      `json:"name"`                // 菜单名称
	Meta      MenuMeta    `json:"meta"`                // 菜单信息
	Method    string      `json:"method"`              //
	Chosen    bool        `json:"is_choose,omitempty"` //
	Redirect  string      `json:"redirect,omitempty"`  // 重定向路径 Path + Children[0].Path
	Children  []MenuModel `json:"children,omitempty"`  // 子级菜单
	Component string      `json:"component,omitempty"` // 对应vue中的map name
	Hidden    bool        `json:"hidden,omitempty"`    // 是否隐藏
	MenuType  uint8       `json:"-"`                   //
}

func (m Menu) Iterate() linq.Iterator {
	src := reflect.ValueOf(m)
	l := src.NumField()
	index := 0

	return func() (item interface{}, ok bool) {
		ok = index < l
		if ok {
			item = src.Field(index).Interface()
			index++
		}

		return
	}

}

func setMenu(menus []Menu, parentID string, hasApi bool) (out []MenuModel) {
	var menuArr []Menu
	linq.From(menus).Where(func(c interface{}) bool {
		return c.(Menu).ParentID == parentID
	}).OrderBy(func(c interface{}) interface{} {
		return c.(Menu).Sequence
	}).ToSlice(&menuArr)
	if len(menuArr) == 0 {
		return
	}
	for _, item := range menuArr {
		if (item.MenuType == MenuTypeOperate && !hasApi) || item.Status == 2 {
			// 操作菜单不显示
			continue
		}
		menu := MenuModel{
			MenuId:   item.MenuId,
			Path:     item.Path,
			Chosen:   item.Chosen,
			Name:     item.Code,
			Method:   item.Method,
			Meta:     MenuMeta{Title: item.Name, Icon: item.Icon, File: item.File},
			MenuType: item.MenuType,
			Children: []MenuModel{}}
		//查询是否有子级
		menuChildren := setMenu(menus, item.MenuId, hasApi)

		if len(menuChildren) > 0 {
			// 新增菜单后 父级被勾选 子级未勾选
			for _, mc := range menuChildren {
				if mc.Chosen == false {
					menu.Chosen = false
					break
				}
			}
			if menu.MenuType == MenuTypeModule {
				menu.Redirect = path.Join(menu.Path, menuChildren[0].Path)
			}
			menu.Children = menuChildren
		}
		out = append(out, menu)
	}
	return
}

func ReadMenuJson(filename string) []MenuModel {
	data, e := ioutil.ReadFile(filename)
	if e != nil {
		log.Fatalln(e)
	}

	var menu Menu
	e = json.Unmarshal(data, &menu)
	if e != nil {
		log.Fatalln(e)
	}

	var menus []Menu
	splitMenus(menu, &menus)
	if len(menus) <= 0 {
		return nil
	}
	topMenu := menus[0]
	menus = menus[1:]
	return setMenu(menus, topMenu.MenuId, false)
}

func splitMenus(menu Menu, menus *[]Menu) {
	subMenus := menu.Children
	menu.Children = nil
	*menus = append(*menus, menu)
	for _, m := range subMenus {
		splitMenus(m, menus)
	}
	return
}
