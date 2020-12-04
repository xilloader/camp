package rbac

import "github.com/xilloader/camp/model"

const (
	// 一个系统通常只有一个Landlord级别的租户
	TenantLevelIsLandlord = "Landlord"
	// Landlord 分配的租户
	// 可通过角色类型分配角色在租户下的权限
	TenantLevelIsTenant = "Tenant"
	// Tenant 分配的租户以及 SubTenant 不可分配租户
	// 多个SubTenant中若存在权限相同或相似的角色较多,可不分配子租户,清晰角色类型即可
	TenantLevelIsSubTenant = "SubTenant"
)

type Tenant struct {
	model.IoModel
	TenantId string `json:"tenant_id" gorm:"primary_key;type:varchar(128);default:'';comment:'租户id'"`
	FatherId string `json:"father_id" gorm:"type:varchar(128);default:'';comment:'父级租户'"`
	Name     string `json:"name" gorm:"type:varchar(128);index:idx_name;uniqueIndex:idx_name;default:'';comment:'名称'"`
	Level    string `json:"level" gorm:"type:varchar(128);default:'';comment:'租户等级 landlord tenant subTenant'"`
	Status   uint8  `json:"status" gorm:"type:tinyint(1);default:1;comment:'状态(0:未激活 1:正常 2:暂停使用)'"`
	Memo     string `json:"memo" gorm:"type:varchar(256);default:'';comment:'备注'" json:"memo" form:"memo"`
}
