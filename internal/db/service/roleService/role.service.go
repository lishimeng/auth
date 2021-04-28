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
