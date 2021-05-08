package authRoleApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/repo"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/auth/internal/db/service/userRolesService"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/auth/internal/respcode"
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

type AuthRoleReq struct {
	Uid int `json:"uid"`
	Rid int `json:"rid"`
}

type AuthRoleResp struct {
	app.Response
}

func Add(ctx iris.Context) {
	log.Info("add role")

	var err error
	orgId := common.GetOrg(ctx)

	var req AuthRoleReq
	var resp AuthRoleResp

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read param failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	var ur = model.AuthUserRoles{
		UserId:    req.Uid,
		RoleId:    req.Rid,
		OrgId:     orgId,
	}
	err = userRolesService.AddUserRole(&ur)
	if err != nil {
		log.Info("add role failed")
		log.Info(err)
		resp.Code = respcode.AddUserFailed
		common.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
	return
}

func Del(ctx iris.Context) {
	log.Info("del role")

	var err error
	var resp app.Response
	orgId := common.GetOrg(ctx)
	id := ctx.Params().GetIntDefault("id", 0)
	if id <= 0 {
		log.Info("id nil")
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}

	aur, err := repo.GetAuthUserRoleById(*app.GetOrm(), id)
	if err != nil {
		log.Info("not found aur:%d", id)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	if aur.OrgId != orgId {
		log.Info("org not match:%d:%d", aur.OrgId, orgId)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}

	err = repo.DelAuthUserRole(*app.GetOrm(), aur)
	if err != nil {
		log.Info("del aur failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func DelUserRole(ctx iris.Context) {
	log.Info("add user role")

	var err error
	var resp app.Response
	orgId := common.GetOrg(ctx)
	uid := ctx.Params().GetIntDefault("uid", 0)
	rid := ctx.Params().GetIntDefault("rid", 0)
	if uid <= 0 {
		log.Info("uid nil")
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	if rid <= 0 {
		log.Info("rid nil")
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}

	aur, err := repo.GetAuthUserRoleByUserAndRole(*app.GetOrm(), uid, rid)
	if err != nil {
		log.Info("not found aur:%d:rid:%d", uid, rid)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	if aur.OrgId != orgId {
		log.Info("org not match:%d:%d", aur.OrgId, orgId)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}

	err = repo.DelAuthUserRole(*app.GetOrm(), aur)
	if err != nil {
		log.Info("del aur failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
