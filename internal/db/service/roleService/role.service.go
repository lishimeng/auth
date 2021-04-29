package roleService

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/repo"
)

func GetUserRoles(uid int) (roles []model.AuthUserRoles, err error) {
	ctx := app.GetOrm()
	roles, err = repo.GetAuthUserRolesByUser(*ctx, uid)
	return
}

func GetOrgRoles(oid int) (roles []model.AuthRoleOrganization, err error) {
	ctx := app.GetOrm()
	roles, err = repo.GetAuthRolesByOrg(*ctx, oid)
	return
}

func GetRole(rid int) (r model.AuthRole, err error) {
	ctx := app.GetOrm()
	r, err = repo.GetAuthRoleById(*ctx, rid)
	return
}
