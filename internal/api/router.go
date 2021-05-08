package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/lishimeng/auth/internal/api/authRoleApi"
	"github.com/lishimeng/auth/internal/api/authUserApi"
	"github.com/lishimeng/auth/internal/api/tokenApi"
	"github.com/lishimeng/auth/internal/api/userApi"
)

func Route(app *iris.Application) {
	app.Use(recover.New())
	route(app.Party("/api"))
}

func route(root iris.Party) {
	token(root.Party("/token"))
	user(root.Party("/user"))
	authRoles(root.Party("/authRoles"))
	authUser(root.Party("/authUser"))
}

func token(p iris.Party) {
	p.Get("/verify", tokenApi.VerifyToken)
}

func user(p iris.Party) {
	p.Post("/sign_in", userApi.SignIn)
	p.Post("/sign_in_card", userApi.SignInCard)
	p.Post("/logout", Authorization, userApi.Logout)
	p.Get("/info/{id}", userApi.GenUserInfo)

	p.Post("/password/change", userApi.ChangePassword)
	p.Post("/password/change", userApi.ChangePasswordWithKey)
	p.Post("/password/reset", userApi.ResetPassword)
}

// authUser
func authUser(p iris.Party) {
	p.Post("/add", authUserApi.Add)
	p.Get("/", Authorization, authUserApi.GetUserList)
	p.Get("/{id}", authUserApi.GetUserInfo)
	p.Put("/{id}", authUserApi.UpdateUserInfo)
	p.Put("/{id}/status", authUserApi.UpdateUserStatus)
	p.Put("/roles/change/{id}", authUserApi.UpdateUserRoles)  // deprecate
}

// authRoles
func authRoles(p iris.Party) {
	p.Get("/", authRoleApi.GetRoleList) // 角色列表
	p.Post("/", Authorization, authRoleApi.Add) // 添加角色
	p.Delete("/{id}", Authorization, authRoleApi.Del) // 删除角色(id:角色关系表id)
	p.Delete("/{uid}/{rid}", Authorization, authRoleApi.DelUserRole) // 删除用户的角色,需要通过user和role查询
}
