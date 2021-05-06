package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/auth/internal/api/tokenApi"
	"github.com/lishimeng/auth/internal/api/userApi"
	"github.com/lishimeng/auth/internal/authRoleApi"
	"github.com/lishimeng/auth/internal/authUserApi"
)

func Route(app *iris.Application) {
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
	p.Post("/logout", userApi.Logout)
	p.Get("/info/{id}", userApi.GenUserInfo)

	p.Post("/password/change", userApi.ChangePassword)
	p.Post("/password/change", userApi.ChangePassword)
	p.Post("/password/reset", userApi.ResetPassword)
}

// authUser
func authUser(p iris.Party) {
	p.Post("/add", authUserApi.Add)
	p.Get("/", authUserApi.GetUserList)
	p.Get("/{id}", authUserApi.GetUserInfo)
	p.Put("/{id}", authUserApi.UpdateUserInfo)
	p.Put("/status/change/{id}", authUserApi.UpdateUserStatus)
	p.Put("/roles/change/{id}", authUserApi.UpdateUserRoles)
}

// authRoles
func authRoles(p iris.Party) {
	p.Get("/", authRoleApi.GetRoleList)
}
