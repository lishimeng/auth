package userService

import (
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/repo"
	"github.com/lishimeng/auth/internal/password"
	persistence "github.com/lishimeng/go-orm"
)

func AddUser(u *model.AuthUser) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// save user
		u.Status = 1 // TODO
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
		e = repo.UpdateUserPassword(ctx, u)
		return e
	})

	return
}

func GetUser(uid int) (u model.AuthUser, err error) {
	ctx := app.GetOrm()
	u, err = repo.GetAuthUserById(*ctx, uid)
	return
}

func GetAuthUserOrg(u model.AuthUser) (uo model.AuthUserOrganization, err error) {
	ctx := app.GetOrm()
	uo, err = repo.GetAuthUserOrg(*ctx, u)
	return
}

// 修改用户状态
func UpdateUserStatusById(id, status int) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// 获取要修改的用户
		u, e := GetUser(id)
		if e != nil {
			return
		}

		// status 列
		var cols []string
		if status > repo.ConditionIgnore {
			u.Status = status
			cols = append(cols, "Status")
		}

		// 修改的时间
		u.UpdateTime = time.Now()
		cols = append(cols, "UpdateTime")

		// 修改状态
		e = repo.UpdateUserStatus(u, cols...)
		return
	})
	return
}

// 修改用户信息
func UpdateUserById(au model.AuthUser) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		u, e := GetUser(au.Id)
		if e != nil {
			return
		}

		// columns
		var cols []string
		if len(au.UserName) > 0 {
			u.UserName = au.UserName
			cols = append(cols, "UserName")
		}
		if len(au.UserNo) > 0 {
			u.UserNo = au.UserNo
			cols = append(cols, "UserNo")
		}
		if len(au.Email) > 0 {
			u.Email = au.Email
			cols = append(cols, "Email")
		}
		if len(au.Phone) > 0 {
			u.Phone = au.Phone
			cols = append(cols, "Phone")
		}
		if au.Status > 0 {
			u.Status = au.Status
			cols = append(cols, "Status")
		}

		// 修改的时间
		u.UpdateTime = time.Now()
		cols = append(cols, "UpdateTime")

		// 修改用户信息
		e = repo.UpdateUserInfo(u, cols...)
		return
	})
	return
}
func ChangePassword(u model.AuthUser, passwd string) (err error) {
	ctx := app.GetOrm()
	u.Password = passwd
	err = repo.UpdateUserPassword(*ctx, &u)
	return
}
