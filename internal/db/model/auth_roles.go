package model

type AuthRole struct {
	Pk
	RoleDescription string `orm:"column(role_desc)"`
	RoleName        string `orm:"column(role_name)"`
	TableChangeInfo
}
