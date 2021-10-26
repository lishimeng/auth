package authUserApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/go-log"
)

type AddReq struct {
	UserNo   string `json:"userNo,omitempty"`
	UserName string `json:"userName,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status,omitempty"`
}

type AddResp struct {
	Uid int `json:"uid,omitempty"`
	app.Response
}

// 新增用户
func Add(ctx iris.Context) {
	log.Debug("add user")
	var req AddReq
	var resp AddResp
	var err error

	var tok jwt.Claims
	common.GetCtxToken(ctx, &tok)

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("req err")
		resp.Code = respcode.AddUserFailed
		resp.Message = "Add User Failed"
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

	// 新增、get_new_uid
	_, err = userService.AddUser(&u)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		resp.Message = "Add User Failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	auo := model.AuthUserOrganization{
		UserId: u.Id,
		OrgId:  tok.OID,
	}

	err = userService.AddUserOrg(&auo)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		resp.Message = "Add User Failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Uid = u.Id
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
