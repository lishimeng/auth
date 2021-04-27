package model

type AuthRoleOrganization struct {
	Pk
	RoleId int `orm:"column(role_id);unique"`
	OrgId  int `orm:"column(org_id)"`
	TableInfo
}
