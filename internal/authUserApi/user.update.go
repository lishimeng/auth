package authUserApi

import (
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/service/userRolesService"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/go-log"
)

// type UserRolesReq struct {
// 	UserId  int    `json:"userId,omitempty"`
// 	RoleIds string `json:"roleIds,omitempty"`
// }

// 修改用户状态
func UpdateUserStatus(ctx iris.Context) {
	log.Debug("update user status")
	var resp app.Response

	// userId、status(new)
	uid := ctx.Params().GetIntDefault("id", 0)
	status := ctx.URLParamIntDefault("status", 0)

	// check
	if uid == 0 {
		log.Debug("uid nil")
		resp.Code = -1
		resp.Message = "uid nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if status == 0 {
		log.Debug("status nil")
		resp.Code = -1
		resp.Message = "status nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	// 修改用户状态
	e := userService.UpdateUserStatusById(uid, status)
	if e != nil {
		log.Debug("can't update user status")
		resp.Code = -1
		resp.Message = "create update fail"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("update user status success, id:%d", uid)
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

// 修改用户信息
func UpdateUserInfo(ctx iris.Context) {
	log.Debug("update user")
	var req model.AuthUser
	var resp app.PagerResponse

	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Debug("req err")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "req err"
		common.ResponseJSON(ctx, resp)
		return
	}

	// check
	if req.Id == 0 {
		log.Debug("param id nil")
		resp.Code = -1
		resp.Message = "param id nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.UserNo) == 0 {
		log.Debug("param userNo nil")
		resp.Code = -1
		resp.Message = "param userNo nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.UserName) == 0 {
		log.Debug("param userName nil")
		resp.Code = -1
		resp.Message = "param userName nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Email) == 0 {
		log.Debug("param email nil")
		resp.Code = -1
		resp.Message = "param email nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	//service.修改用户信息
	e := userService.UpdateUserById(req)
	if e != nil {
		log.Debug("can't update user")
		resp.Code = -1
		resp.Message = "create update fail"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("update user success, id:%d", req.Id)
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

// 修改用户角色
func UpdateUserRoles(ctx iris.Context) {
	log.Debug("update user roles")
	var resp app.Response

	// userId
	uid := ctx.Params().GetIntDefault("id", 0)
	// String roleIds
	roleIds := ctx.URLParamDefault("roleIds", "")

	// check
	if len(roleIds) == 0 {
		log.Debug("roleIds nil")
		resp.Code = -1
		resp.Message = "roleIds nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	var aur model.AuthUserRoles
	aur.UserId = uid
	// 删除用户的所有角色
	e := userRolesService.DeleteUserRoles(aur)
	if e != nil {
		log.Debug("can't delete user role")
		resp.Code = -1
		resp.Message = "create delete fail"
		common.ResponseJSON(ctx, resp)
		return
	}
	// 获取 roleId 列表： strings --> list[]
	roleList := strings.Split(roleIds, ",")
	for role := range roleList {
		var ur model.AuthUserRoles
		ur.RoleId = role
		ur.UserId = uid
		ur.CreateTime = time.Now()
		// 新增 user_role
		userRolesService.AddUserRole(ur)
	}

	log.Debug("update user roles success, id:%d", uid)
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
