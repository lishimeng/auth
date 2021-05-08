package authUserApi

import (
	"strconv"
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

type UserRolesReq struct {
	RoleIds string `json:"roleIds,omitempty"`
}

type StatusReq struct {
	Status int `json:"status"`
}

// UpdateUserStatus 修改用户状态
func UpdateUserStatus(ctx iris.Context) {
	log.Debug("update user status")
	var req StatusReq
	var resp app.Response

	// userId、status(new)
	uid := ctx.Params().GetIntDefault("id", 0)
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Debug("req err")
		log.Debug(err)
		resp.Code = common.RespCodeInternalError
		resp.Message = "req err"
		common.ResponseJSON(ctx, resp)
		return
	}
	// check
	if uid == 0 {
		log.Debug("uid nil")
		resp.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}
	if req.Status == 0 {
		log.Debug("status nil")
		resp.Code = common.RespCodeSuccess // TODO ignore, if no param:status
		resp.Message = "status nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	// 修改用户状态
	e := userService.UpdateUserStatus(uid, req.Status)
	if e != nil {
		log.Debug("can't update user status")
		resp.Code = common.RespCodeInternalError
		resp.Message = "update status fail"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("update user status success, id:%d", uid)
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(ctx iris.Context) {
	log.Debug("update user")
	var req model.AuthUser
	var resp app.PagerResponse

	uid := ctx.Params().GetIntDefault("id", 0)
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Debug("req err")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "req err"
		common.ResponseJSON(ctx, resp)
		return
	}

	//service.修改用户信息
	req.Id = uid
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

// UpdateUserRoles 修改用户角色
//
// @deprecate
func UpdateUserRoles(ctx iris.Context) {
	log.Debug("update user roles")
	var req UserRolesReq
	var resp app.Response

	// userId
	uid := ctx.Params().GetIntDefault("id", 0)
	// String roleIds
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Debug("req err")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "req err"
		common.ResponseJSON(ctx, resp)
		return
	}

	// 删除用户的所有角色
	var aur model.AuthUserRoles
	aur.UserId = uid
	log.Debug("delete user role, uid:%d", uid)
	e := userRolesService.DeleteUserRoles(aur)
	if e != nil {
		log.Debug("can't delete user role")
		resp.Code = -1
		resp.Message = "create delete fail"
		common.ResponseJSON(ctx, resp)
		return
	}
	// 获取 roleId 列表： strings --> list[]
	roleList := strings.Split(req.RoleIds, ",")
	log.Debug(roleList)
	for _, role := range roleList {
		var ur model.AuthUserRoles
		ur.RoleId, err = strconv.Atoi(role)
		if err != nil {
			log.Debug("can't transfer string to int")
		}
		if ur.RoleId != 0 {
			ur.UserId = uid
			ur.CreateTime = time.Now()
			// add user_role
			_ = userRolesService.AddUserRole(&ur)
		}
	}

	log.Debug("update user roles success, id:%d", uid)
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
