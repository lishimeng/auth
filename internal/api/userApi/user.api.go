package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/go-log"
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

func GenUserInfo(ctx iris.Context) {
	var err error
	var resp Resp
	uid := ctx.Params().GetIntDefault("id", 0)
	if uid <= 0 {
		log.Info("uid nil")
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	// 校验用户
	u, err := userService.GetUser(uid)
	if err != nil {
		log.Info("no user:%d", uid)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
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

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)

	return
}

func Logout(ctx iris.Context) {
	ctx.GetHeader("")
	var resp app.Response
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}