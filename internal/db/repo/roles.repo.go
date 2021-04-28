package repo

import (
	"github.com/lishimeng/auth/internal/db/model"
	persistence "github.com/lishimeng/go-orm"
)

func GetAuthUserRolesByUser(ctx persistence.OrmContext, uid int) (aur []model.AuthUserRoles, err error) {
	_, err = ctx.Context.QueryTable(new(model.AuthUserRoles)).Filter("UserId", uid).All(&aur)
	return
}
