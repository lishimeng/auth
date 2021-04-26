package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/auth/internal/api/tokenApi"
)

func Route(app *iris.Application) {
	p := app.Party("/tokenApi")
	token(p)
}

func token(p iris.Party) {
	p.Post("/gen", tokenApi.GenToken)
	p.Get("/verify", tokenApi.VerifyToken)
}
