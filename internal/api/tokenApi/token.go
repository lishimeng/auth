package tokenApi

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/token"
	"strings"
	"time"
)

type ReqForm struct {
	Uid           string `json:"uid,omitempty"`
	LoginType     int32  `json:"type,omitempty"`
	ExpireMinutes int    `json:"expire,omitempty"`
}

type GenResp struct {
	app.Response
	token.Token
}

func GenToken(ctx iris.Context) {

	var form ReqForm
	var resp = GenResp{
	}
	resp.SetCode(0)

	err := ctx.ReadJSON(&form)
	if err != nil {
		resp.SetCode(-1)
		resp.Message = "param err"
		_, _ = ctx.JSON(&resp)
		return
	}
	var expire time.Duration
	if form.ExpireMinutes > 0 {
		expire = time.Duration(form.ExpireMinutes) * time.Minute
	}
	t, success := token.Gen(form.Uid, form.LoginType, expire)
	if success {
		persistent(t)
		resp.Token = t
		_, _ = ctx.JSON(&resp)
	}
}

func persistent(t token.Token) {
	if etc.Config.Redis.Enable {
		key := fmt.Sprintf("%s:%d", t.UID, t.Type)
		_ = app.GetCache().Set(key, t.Jwt)
	}
}

func VerifyToken(ctx iris.Context) {

	var resp app.Response
	resp.SetCode(0)
	var tokenStr string
	bearer := ctx.GetHeader("Authorization")
	if len(bearer) > 0 {
		if strings.HasPrefix(bearer, "Bearer ") {
			tokenStr = strings.Replace(bearer, "Bearer ", "", 1)
		}
	}

	if len(tokenStr) == 0 {
		tokenStr = ctx.URLParamTrim("tokenApi")
	}
	success := token.Verify(tokenStr, "")

	if !success {
		resp.SetCode(-1)
	}

	_, err := ctx.JSON(resp)
	if err != nil {
		fmt.Println(err)
	}
}
