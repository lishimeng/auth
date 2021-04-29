package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/respcode"
)

type AddReq struct {
	UserNo   string `json:"user_no"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Phone    string `json:"mobile_phone"`
	Password string `json:"password"`
}

type AddResp struct {
	app.Response
	UserNo   string `json:"user_no"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Phone    string `json:"mobile_phone"`
	Password string `json:"password"`
}

func Add(ctx iris.Context) {
	var req AddReq
	var resp AddResp
	var err error
	u := model.AuthUser{
		UserNo:          req.UserNo,
		UserName:        req.UserName,
		Email:           req.Email,
		Phone:           req.Phone,
		Password:        req.Password,
	}

	err = userService.AddUser(&u)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		common.ResponseJSON(ctx, resp)
		return
	}

	// -------------------------------------
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
