package rbac

import "github.com/xilloader/camp/model"

type Role struct {
	model.IoModel
	RoleId   string `json:"role_id" gorm:"primary_key;type:varchar(64);default:'';comment:'角色id'"`
	TenantId string `json:"tenant_id" gorm:"type:varchar(128);index:idx_tenant_id_name;default:'';comment:'租戶id'"`// 角色只在该租户下可用
	ParentID string `json:"parent_id" gorm:"type:varchar(64);default:'';comment:'角色父级id'"`
	Name     string `json:"name" gorm:"type:varchar(128);index:idx_tenant_id_name;default:'';comment:'名称'"`
	RoleType uint8  `json:"role_type" gorm:"type:int;default:0;comment:'角色类型'"` //角色类型与菜单领域有关
	Sequence uint64 `json:"sequence" gorm:"type:int;default:0;comment:'排序值'"`
	IsSuper  uint8  `json:"is_super" gorm:"type:tinyint(1);default:0;comment:'是否是超管 0不是 1是'"` // 超管角色通常不可随意修改, 可统一管理
	Memo     string `json:"memo" gorm:"type:varchar(256);default:'';comment:'备注 不能超过60个字符'"`
}
