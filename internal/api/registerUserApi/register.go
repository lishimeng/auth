package registerUserApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/go-log"
)

type RegisterReq struct {
	UserNo   string `json:"userNo,omitempty"`
	UserName string `json:"userName,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status,omitempty"`
	OrgId    int    `json:"orgId,omitempty"`
	Code     string `json:"code,omitempty"`
	ImgCode  string `json:"imgCode,omitempty"`
}

type RegisterResp struct {
	Uid int `json:"uid,omitempty"`
	app.Response
}

func Register(ctx iris.Context) {
	log.Debug("Register user")
	var req RegisterReq
	var resp RegisterResp
	var err error

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("Wrong request parameters.")
		resp.Code = respcode.AddUserFailed
		resp.Message = "Wrong request parameters."
		common.ResponseJSON(ctx, resp)
		return
	}
	if !(req.OrgId > 0 && len(req.UserNo) > 0 &&
		len(req.UserName) > 0 && len(req.Email) > 0 &&
		len(req.Password) > 0) {
		log.Info("Parameter is missing.")
		resp.Code = respcode.AddUserFailed
		resp.Message = "Parameter is missing."
		common.ResponseJSON(ctx, resp)
		return
	}
	// 判断 邮件验证码
	log.Info("mail: %s, code: %s", req.Email,req.Code)
	isCode, err := ModifyMailCode(req.Email, req.Code)
	if err != nil {
		log.Info("Mail verification code has expired.")
		log.Info(err)
		resp.Code = respcode.AddUserFailed
		resp.Message = "Mail verification code has expired."
		common.ResponseJSON(ctx, resp)
		return
	}
	if !isCode {
		log.Info("Mail verification code has expired.")
		resp.Code = respcode.AddUserFailed
		resp.Message = "Mail verification code has expired."
		common.ResponseJSON(ctx, resp)
		return
	}


	// TODO 判断 图片验证码

	ut := model.TableChangeInfo{
		Status: common.UserStatusActivate,
	}
	u := model.AuthUser{
		UserNo:          req.UserNo,
		UserName:        req.UserName,
		Email:           req.Email,
		Phone:           req.Phone,
		Password:        req.Password,
		TableChangeInfo: ut,
	}

	// 新增 get_new_uid
	_, err = userService.AddUser(&u)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		resp.Message = "Register Failed. User ID or Email already exists."
		common.ResponseJSON(ctx, resp)
		return
	}

	auo := model.AuthUserOrganization{
		UserId: u.Id,
		OrgId:  req.OrgId,
	}

	err = userService.AddUserOrg(&auo)
	if err != nil {
		resp.Code = respcode.AddUserFailed
		resp.Message = "Register Failed."
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Uid = u.Id
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func ModifyMailCode(email,code string) (isCode bool, err error) {
	var cacheCode string
	err = app.GetCache().Get(email,&cacheCode)
	if err != nil {
		return
	}
	if code != cacheCode{
		log.Info("cache code: %s", cacheCode)
		return
	}
	return true, err
}
