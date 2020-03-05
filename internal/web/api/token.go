package api

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/lishimeng/auth/internal/token"
	"strings"
)

func TokenApi(app *iris.Application) {

	p := app.Party("token")
	p.Post("/gen", genToken)
	p.Get("/verify", verifyToken)
}

type TokenReqForm struct {
	Uid string `form:"uid"`
	Password string `form:"password"`
	LoginType int32 `form:"loginType"`
}

type TokenGenResp struct {
	token.Token
	Response
}

func genToken(ctx iris.Context) {

	var form TokenReqForm
	var resp = TokenGenResp{
		Response: Response{Code: 0},
	}
	err := ctx.ReadForm(&form)
	if err != nil {
		resp.Response.Code = -1
		return
	}

	// TODO login

	//
	t, success := token.Gen(form.Uid, form.LoginType)
	if success {
		resp.Token = t
		_, err = ctx.JSON(&resp)
	}
}

func verifyToken(ctx iris.Context) {

	var resp = Response{
		Code:    0,
	}
	var tokenStr string
	bearer := ctx.GetHeader("Authorization")
	if len(bearer) > 0 {
		if strings.HasPrefix(bearer, "Bearer ") {
			tokenStr = strings.Replace(bearer, "Bearer ", "", 1)
		}
	}

	if len(tokenStr) == 0 {
		tokenStr = ctx.URLParamTrim("token")
	}
	success := token.Verify(tokenStr, "")

	if !success {
		resp.Code = -1
	}

	_, err := ctx.JSON(resp)
	if err != nil {
		fmt.Println(err)
	}
}