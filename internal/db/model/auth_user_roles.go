package model

type AuthUserRoles struct {
	Pk
	UserId int `orm:"column(user_id)"`
	RoleId int `orm:"column(role_id)"`
	OrgId  int `orm:"column(org_id)"`
	TableInfo
}

func (t *AuthUserRoles) TableUnique() [][]string {
	return [][]string{
		{"UserId", "RoleId"},
	}
}
