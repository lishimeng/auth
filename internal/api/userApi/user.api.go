package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
)

type Req struct {
	Uid      int `json:"uid,omitempty"`
	LoginName string `json:"loginName,omitempty"`
	Password string `json:"password,omitempty"`
}

type Resp struct {
	app.Response
	Jwt string `json:"jwt,omitempty"`
	OrgId int `json:"orgId,omitempty"`
	Roles []int `json:"roles,omitempty"`
}

func Logout(ctx iris.Context) {
	ctx.GetHeader("")
	var resp app.Response
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}