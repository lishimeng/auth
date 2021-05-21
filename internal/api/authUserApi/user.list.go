package authUserApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/repo"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/go-log"
)

type RoleInfo struct {
	RoleId   int    `json:"roleId,omitempty"`
	RoleName string `json:"roleName,omitempty"`
}

type UserInfo struct {
	UserNo   string `json:"userNo,omitempty"`
	UserName string `json:"userName,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Status   int    `json:"status,omitempty"`
	UserId   int    `json:"userId,omitempty"`
}

type UserInfoResp struct {
	app.Response
	UserInfo
	Roles []int `json:"roles,omitempty"`
}

// 获取用户列表
func GetUserList(ctx iris.Context) {
	log.Debug("get user list")
	var resp app.PagerResponse
	var pageSize = ctx.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	var tok jwt.Claims
	common.GetCtxToken(ctx, &tok)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNo,
	}

	// org_users
	page, auos, err := repo.GetOrgUsers(tok.OID, page)
	if err != nil {
		log.Debug("get org users failed oid:%d", tok.OID)
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get org users failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	// user_list
	for _, auo := range auos {
		u, err := userService.GetUser(auo.UserId)
		if err != nil {
			log.Debug("get org users failed uid:%d", auo.UserId)
			log.Info(err)
		} else {
			var userInfo = UserInfo{
				UserNo:   u.UserNo,
				UserName: u.UserName,
				Email:    u.Email,
				Phone:    u.Phone,
				Status:   u.Status,
				UserId:   u.Id,
			}
			page.Data = append(page.Data, userInfo)
		}
	}

	resp.Pager = page
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

// 获取用户信息
func GetUserInfo(ctx iris.Context) {
	var err error
	var resp UserInfoResp
	// userId
	uid := ctx.Params().GetIntDefault("id", 0)
	if uid <= 0 {
		log.Info("uid nil")
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	// user
	u, err := userService.GetUser(uid)
	if err != nil {
		log.Info("no user:%d", uid)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	resp.UserId = u.Id
	resp.UserName = u.UserName
	resp.UserNo = u.UserNo
	resp.Status = u.Status
	resp.Phone = u.Phone
	resp.Email = u.Email

	// role_list
	aurs, err := roleService.GetUserRoles(u.Id)
	if err != nil {
		log.Info("get auth role fail uid:%d", u.Id)
		log.Info(err)
	} else {
		for _, r := range aurs {
			var role int
			role = r.RoleId
			resp.Roles = append(resp.Roles, role)
		}
	}
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
