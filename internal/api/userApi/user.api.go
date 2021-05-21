package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/go-log"
)

type Req struct {
	Uid       int    `json:"uid,omitempty"`
	LoginName string `json:"loginName,omitempty"`
	Password  string `json:"password,omitempty"`
}

type Resp struct {
	app.Response
	Jwt   string `json:"jwt,omitempty"`
	Uid   int    `json:"uid,omitempty"`
	OrgId int    `json:"orgId,omitempty"`
	Roles []int  `json:"roles,omitempty"`
}

func Logout(ctx iris.Context) {
	var tok jwt.Claims
	common.GetCtxToken(ctx, &tok)
	log.Info("logout uid:%d, oid:%d", tok.UID, tok.OID)
	// TODO
	var resp app.Response
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
