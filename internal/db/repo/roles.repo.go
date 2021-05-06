package repo

import (
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/auth/internal/db/model"
	persistence "github.com/lishimeng/go-orm"
)

func GetAuthUserRolesByUser(ctx persistence.OrmContext, uid int) (aur []model.AuthUserRoles, err error) {
	_, err = ctx.Context.QueryTable(new(model.AuthUserRoles)).Filter("UserId", uid).All(&aur)
	return
}

func GetAuthRolesByOrg(ctx persistence.OrmContext, oid int) (aro []model.AuthRoleOrganization, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("OrgId", oid)
	_, err = ctx.Context.QueryTable(new(model.AuthRoleOrganization)).SetCond(cond.AndCond(cond1)).All(&aro)
	return
}

func GetAuthRoleById(ctx persistence.OrmContext, rid int) (r model.AuthRole, err error) {
	r.Id = rid
	err = ctx.Context.Read(&r)
	return
}

func DelAuthUserRole(ctx persistence.OrmContext, aur model.AuthUserRoles) (err error) {

	_, err = ctx.Context.Delete(aur)

	return
}
