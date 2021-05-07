package tokenApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/jwt"
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

	c, success := common.Authorization(ctx)

	if !success {
		resp.SetCode(common.RespCodeNeedAuth)
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Claims = *c
	resp.SetCode(common.RespCodeSuccess)
	common.ResponseJSON(ctx, resp)
}
