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
		resp.Code = respcode.SignInFailed
		common.ResponseJSON(ctx, resp)
		return
	}
	auo, err := userService.GetAuthUserOrg(u)
	if err != nil {
		log.Info("no auth org info")
		log.Info(err)
	} else {
		resp.OrgId = auo.OrgId
	}
	// 创建jwt
	t, success := token.Gen(u.PkString(), 1, 0)
	if !success {
		log.Info("create jwt failed")
		resp.Code = respcode.SignInFailed
		common.ResponseJSON(ctx, resp)
		return
	}

	// role_list
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
	// jwt 存入cache
	err = app.GetCache().Set(t.Jwt, t.UID)
	if err != nil {
		log.Info(err)
	}
	// TODO 返回jwt

	resp.Code = common.RespCodeSuccess
	resp.Jwt = t.Jwt
	common.ResponseJSON(ctx, resp)

}