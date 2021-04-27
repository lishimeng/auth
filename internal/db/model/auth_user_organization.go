package model

type AuthUserOrganization struct {
	Pk
	UserId int `orm:"column(user_id);unique"`
	OrgId  int `orm:"column(org_id)"`
	TableInfo
}
