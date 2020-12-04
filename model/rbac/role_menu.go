package rbac

type M2mRoleMenu struct {
	RRoleId string `json:"r_role_id" gorm:"primary_key;type:varchar(128);default:'';comment:'Role.RoleId'"`
	Roles   []Role `json:"roles,omitempty" gorm:"foreignKey:RoleId;references:RRoleId"`

	MMenuId string `json:"m_menu_id" gorm:"primary_key;type:varchar(128);default:'';comment:'Menu.MenuId'"`
	Menus   []Menu `json:"menus,omitempty" gorm:"foreignKey:MenuId;references:MMenuId"`
}
