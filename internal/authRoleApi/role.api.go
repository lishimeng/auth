package authRoleApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/go-log"
)

// 获取角色列表
func GetRoleList(ctx iris.Context) {
	var resp app.PagerResponse
	oid := ctx.Params().GetIntDefault("id", 0)
	// org_roles
	aros, err := roleService.GetOrgRoles(oid)
	if err != nil {
		log.Info("get org role fail oid:%d", oid)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	// role_list
	for _, r := range aros {
		ar, err := roleService.GetRole(r.RoleId)
		if err != nil {
			log.Info("get auth role fail rid:%d", r.RoleId)
			log.Info(err)
		} else {
			resp.Data = append(resp.Data, ar)
		}
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)

	return
}
