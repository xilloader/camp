package rbac

import "github.com/xilloader/camp/model"

// 用户可属于多个租户
type User struct {
	model.IoModel
	UserId          string `json:"user_id" gorm:"primary_key;type:varchar(128);default:'';comment:'用户id'" `
	Name            string `json:"name" gorm:"type:varchar(128);default:'';comment:'用户名称'"`
	UserName        string `json:"user_name" gorm:"type:varchar(128);default:'';comment:'租户id';unique_index:uk_users_user_name;"`
	Password        string `json:"password" gorm:"type:varchar(128);default:'';comment:'密码'"`
	Status          uint8  `json:"status" gorm:"column:status;type:tinyint(1);default:1;comment:'状态(1:正常 2:未激活 3:暂停使用)'"`
	CreatorTenantId string `json:"creator_tenant_id" gorm:"type:varchar(128);default:'';comment:'用户创建租戶'"` //用户创建租戶
	CreatorUserId   string `json:"creator_user_id" gorm:"type:varchar(128);default:'';comment:'用户创建用户'"`   //用户创建用户
	Memo            string `json:"memo" gorm:"type:varchar(128);default:'';comment:'备注'"`
}
