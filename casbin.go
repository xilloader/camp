package camp

import (
	"github.com/casbin/casbin"
)

const (
	PrefixUserID = "u_"
	PrefixRoleID = "r_"
	ErrKey       = "Casbin"

	casbinModel = `
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) == true \
			&& keyMatch2(r.obj, p.obj) == true \
			&& regexMatch(r.act, p.act) == true \
			|| r.sub == "root"
	`
)

var Enforcer *casbin.Enforcer

// 角色-URL导入
func InitCsbinEnforcer() (err error) {

	var enforcer *casbin.Enforcer

	enforcer, err = casbin.NewEnforcerSafe(casbin.NewModel(casbinModel), true)
	if err != nil {
		return
	}

	Enforcer = enforcer
	return
}

//检查enforcer是否为空
func CheckEnforcer(enforcer *casbin.Enforcer) error {
	if enforcer == nil {
		return NewError("enforcer is nil", ErrKey)
	}
	return nil
}

// 检查用户是否有权限
func CasbinCheckPermission(user, url, methodtype string) (bool, error) {
	return Enforcer.EnforceSafe(PrefixUserID+user, url, methodtype)
}

//重置用户角色
func CasbinAddRoleForUser(user string, roles ...string) error {
	for i := range roles {
		roles[i] = PrefixRoleID + roles[i]
	}
	return enforcerSetRoleForUser(Enforcer, PrefixUserID+user, roles...)
}

//删除用户所有角色
func CasbinDeleteRolesForUser(user string) {
	Enforcer.DeleteRolesForUser(PrefixUserID + user)
}

//为角色分配权限
func CasbinSetRolePermission(reset bool, role string, objs ...string) error {
	act := "GET|POST"
	if !reset {
		return enforcerSetRolePermission(Enforcer, PrefixRoleID+role, act, objs...)
	}
	return enforcerResetRolePermission(Enforcer, PrefixRoleID+role, act, objs...)

}

//删除角色
func CasbinDeleteRole(roles ...string) error {
	for i := range roles {
		roles[i] = PrefixRoleID + roles[i]
	}
	return enforcerDeleteRole(Enforcer, roles...)
}

//重新设置用户角色
func enforcerSetRoleForUser(enforcer *casbin.Enforcer, user string, roles ...string) error {

	if err := CheckEnforcer(enforcer); err != nil {
		return err
	}

	enforcer.DeleteRolesForUser(user)
	for _, role := range roles {
		enforcer.AddRoleForUser(user, role)
	}

	return nil
}

//设置角色权限
func enforcerResetRolePermission(enforcer *casbin.Enforcer, role string, act string, objs ...string) error {

	if err := CheckEnforcer(enforcer); err != nil {
		return err
	}

	//删除角色(或用户)原有权限
	enforcer.DeletePermissionsForUser(role)
	for _, obj := range objs {
		enforcer.AddPermissionForUser(role, obj, act)
	}

	return nil
}

func enforcerSetRolePermission(enforcer *casbin.Enforcer, role string, act string, objs ...string) error {

	if err := CheckEnforcer(enforcer); err != nil {
		return err
	}

	for _, obj := range objs {
		enforcer.AddPermissionForUser(role, obj, act)
	}

	return nil
}

//删除角色及权限
func enforcerDeleteRole(enforcer *casbin.Enforcer, roles ...string) error {

	if err := CheckEnforcer(enforcer); err != nil {
		return err
	}

	for _, role := range roles {
		//删除角色(或用户)权限
		enforcer.DeletePermissionsForUser(role)
		//删除角色
		enforcer.DeleteRole(role)
	}

	return nil
}
