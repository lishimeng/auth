package userService

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/repo"
	"github.com/lishimeng/auth/internal/password"
	persistence "github.com/lishimeng/go-orm"
)

func AddUser(u *model.AuthUser) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// save user
		e = repo.AddUser(ctx, u)
		if e != nil {
			return
		}
		// gen password
		pswEnc, e := password.Generate(*u, u.Password)
		if e != nil {
			return
		}
		// update user
		u.Password = pswEnc
		repo.UpdateUserPassword(ctx, u)
		return e
	})

	return
}
