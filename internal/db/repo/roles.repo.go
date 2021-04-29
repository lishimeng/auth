package repo

import (
	"github.com/lishimeng/auth/internal/db/model"
	persistence "github.com/lishimeng/go-orm"
)

func GetAuthUserRolesByUser(ctx persistence.OrmContext, uid int) (aur []model.AuthUserRoles, err error) {
	_, err = ctx.Context.QueryTable(new(model.AuthUserRoles)).Filter("UserId", uid).All(&aur)
	return
}

func GetAuthRolesByOrg(ctx persistence.OrmContext, oid int) (aro []model.AuthRoleOrganization, err error) {
	_, err = ctx.Context.QueryTable(new(model.AuthRoleOrganization)).Filter("OrgId").All(&aro)
	return
}

func GetAuthRoleById(ctx persistence.OrmContext, rid int) (r model.AuthRole, err error) {
	r.Id = rid
	err = ctx.Context.Read(&r)
	return
}
