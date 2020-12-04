package rbac

type M2mUserRole struct {
	UUserId string `json:"u_user_id" gorm:"primary_key;type:varchar(128);default:'';comment:'User.UserId'"`
	Users   []User `json:"users,omitempty" gorm:"foreignKey:UserId;references:UUserId"`
	RRoleId string `json:"r_role_id" gorm:"primary_key;type:varchar(128);default:'';comment:'Menu.MenuId'"`
	Roles   []Role `json:"roles,omitempty" gorm:"foreignKey:RoleId;references:RRoleId"`
}
