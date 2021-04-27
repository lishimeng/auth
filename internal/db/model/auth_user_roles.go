package model

type AuthUserRoles struct {
	Pk
	UserId int `orm:"column(user_id)"`
	RoleId int `orm:"column(role_id)"`
	TableInfo
}

func (t *AuthUserRoles) TableUnique() [][]string {
	return [][]string{
		{"UserId", "RoleId"},
	}
}
