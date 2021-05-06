package tokenApi

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/auth/internal/token"
	"strings"
)

type ReqForm struct {
	Uid           string `json:"uid,omitempty"`
	LoginType     int32  `json:"type,omitempty"`
	ExpireMinutes int    `json:"expire,omitempty"`
}

type GenResp struct {
	app.Response
	jwt.Claims
}

func VerifyToken(ctx iris.Context) {

	var resp GenResp
	resp.SetCode(0)
	var tokenStr string

	// Authorization Bearer
	bearer := ctx.GetHeader("Authorization")
	if len(bearer) > 0 {
		if strings.HasPrefix(bearer, "Bearer ") {
			tokenStr = strings.Replace(bearer, "Bearer ", "", 1)
		}
	}

	// accessToken:
	// schema://domain/path?accessToken=
	if len(tokenStr) == 0 {
		tokenStr = ctx.URLParamTrim("accessToken")
	}
	c, success := token.Verify(tokenStr)

	if !success {
		resp.SetCode(common.RespCodeNeedAuth)
		return
	}

	resp.Claims = *c
	resp.SetCode(common.RespCodeSuccess)
	_, err := ctx.JSON(resp)
	if err != nil {
		fmt.Println(err)
	}
}
