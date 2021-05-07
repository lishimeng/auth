package authRoleApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/go-log"
)

type RespRoleInfo struct {
	Id              int    `json:"id,omitempty"`
	RoleName        string `json:"roleName,omitempty"`
	RoleDescription string `json:"roleDesc,omitempty"`
	Status          int    `json:"status,omitempty"`
}

// 获取角色列表
func GetRoleList(ctx iris.Context) {
	log.Info("get role list")
	var resp app.PagerResponse
	var tok jwt.Claims
	common.GetCtxToken(ctx, &tok)

	// org_roles
	aros, err := roleService.GetOrgRoles(tok.OID)
	if err != nil {
		log.Info("get org role fail oid:%d", tok.OID)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	// role_list
	for _, r := range aros {
		// 获取 role
		ar, err := roleService.GetRole(r.RoleId)
		if err != nil {
			log.Info("get auth role fail rid:%d", r.RoleId)
			log.Info(err)
		} else {
			var roleInfo = RespRoleInfo{
				Id:              ar.Id,
				RoleName:        ar.RoleName,
				RoleDescription: ar.RoleDescription,
				Status:          ar.Status,
			}
			// 添加到 Data[]
			resp.Data = append(resp.Data, roleInfo)
		}
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
