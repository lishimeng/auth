package authUserApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
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
}

type AddResp struct {
	app.Response
	UserNo   string `json:"userNo,omitempty"`
	UserName string `json:"userName,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"mobilePhone,omitempty"`
	Password string `json:"password,omitempty"`
}

// 新增用户
func Add(ctx iris.Context) {
	log.Debug("add user")
	var req AddReq
	var resp AddResp
	var err error
	err = ctx.ReadJSON(&req)
	log.Info("add user req[]:", req)
	if err != nil {
		log.Info("req err")
		return
	}
	u := model.AuthUser{
		UserNo:   req.UserNo,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}

	// 新增
	err = userService.AddUser(&u)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
