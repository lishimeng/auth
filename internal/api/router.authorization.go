package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
)

func Authorization(ctx iris.Context) {
	c, success := common.Authorization(ctx)
	if !success {
		var resp = app.Response{}
		resp.Code = common.RespCodeNeedAuth
		common.ResponseJSON(ctx, resp)
		return
	}
	common.SaveCtxToken(ctx, c)
	ctx.Next()
}
