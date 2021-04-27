package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/auth/internal/api/tokenApi"
	"github.com/lishimeng/auth/internal/api/userApi"
)

func Route(app *iris.Application) {
	route(app.Party("/api"))
}

func route(root iris.Party) {
	token(root.Party("/token"))
	user(root.Party("/user"))
}

func token(p iris.Party) {
	p.Post("/gen", tokenApi.GenToken)
	p.Get("/verify", tokenApi.VerifyToken)
}

func user(p iris.Party) {
	p.Post("/sign_in", userApi.SignIn)
	p.Post("/logout", userApi.Logout)
}
