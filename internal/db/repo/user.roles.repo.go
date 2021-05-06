package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/db/model"
	persistence "github.com/lishimeng/go-orm"
)

func DeleteUserRole(aur *model.AuthUserRoles, cols ...string) (err error) {
	_, err = app.GetOrm().Context.Delete(aur, cols...)
	return
}

func AddUserRole(ctx persistence.OrmContext, u *model.AuthUserRoles) (err error) {
	_, err = ctx.Context.Insert(u)
	return
}
