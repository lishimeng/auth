package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/auth/internal/token"
	"github.com/lishimeng/go-log"
)

type Req struct {
	LoginName string `json:"loginName,omitempty"`
	Password string `json:"password,omitempty"`
}

type Resp struct {
	app.Response
	Jwt string `json:"jwt,omitempty"`
	Roles []int `json:"roles,omitempty"`
}

func SignIn(ctx iris.Context) {

	var req Req
	var resp Resp
	var err error
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}
	// 校验用户
	u, err := userService.SignIn(req.LoginName, req.Password)
	if err != nil {
		log.Info("sign in failed")
		log.Info(err)
		resp.Code = respcode.RespSignInFailed
		common.ResponseJSON(ctx, resp)
		return
	}
	// 创建jwt
	t, success := token.Gen(u.PkString(), 1, 0)
	if !success {
		log.Info("create jwt failed")
		resp.Code = respcode.RespSignInFailed
		common.ResponseJSON(ctx, resp)
		return
	}

	// TODO role_list
	aurs, err := roleService.GetUserRoles(u.Id)
	if err != nil {
		log.Info("get auth role fail uid:%d", u.Id)
		log.Info(err)
	} else {
		for _, r := range aurs {
			resp.Roles = append(resp.Roles, r.RoleId)
		}
	}
	// TODO role_list 存入cache
	// TODO jwt 存入cache
	// TODO 返回jwt

	resp.Code = common.RespCodeSuccess
	resp.Jwt = t.Jwt
	common.ResponseJSON(ctx, resp)

}

func Logout(ctx iris.Context) {
	ctx.GetHeader("")
	var resp app.Response
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}