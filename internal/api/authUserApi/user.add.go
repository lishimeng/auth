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
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/go-log"
)

type AddReq struct {
	UserNo   string `json:"userNo,omitempty"`
	UserName string `json:"userName,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"mobilePhone,omitempty"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status,omitempty"`
	RoleIds  string `json:"roleIds,omitempty"`
}

// 新增用户
func Add(ctx iris.Context) {
	log.Debug("add user")
	var req AddReq
	var resp app.Response
	var err error

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("req err")
		return
	}

	ut := model.TableChangeInfo{
		Status: req.Status,
	}
	u := model.AuthUser{
		UserNo:          req.UserNo,
		UserName:        req.UserName,
		Email:           req.Email,
		Phone:           req.Phone,
		Password:        req.Password,
		TableChangeInfo: ut,
	}

	// 新增
	err = userService.AddUser(&u)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		common.ResponseJSON(ctx, resp)
		return
	}

	// 获取 roleId 列表： strings --> list[]
	// TODO 删除这部分,不在创建用户时添加权限
	roleList := strings.Split(req.RoleIds, ",")
	log.Info(roleList)
	for _, role := range roleList {
		var ur model.AuthUserRoles
		ur.RoleId, err = strconv.Atoi(role)
		if ur.RoleId != 0 {
			ur.UserId = u.Id
			ur.CreateTime = time.Now()
			// add user_role
			err = userRolesService.AddUserRole(&ur)
		}
		if err != nil {
			log.Info(err)
		}
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
