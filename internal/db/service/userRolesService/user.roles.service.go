package userRolesService

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/repo"
	persistence "github.com/lishimeng/go-orm"
)

func DeleteUserRoles(aur model.AuthUserRoles) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		var cols []string
		cols = append(cols, "UserId")
		e = repo.DeleteUserRole(&aur, cols...)
		return
	})
	return
}

func AddUserRole(u *model.AuthUserRoles) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		e = repo.AddUserRole(ctx, u)
		if e != nil {
			return
		}
		return
	})

	return
}
