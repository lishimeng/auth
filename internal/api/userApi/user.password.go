package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/password"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/go-log"
)

type PasswordReq struct {
	Uid int    `json:"uid"`
	Old string `json:"old"`
	New string `json:"new"`
}

type PasswordResetReq struct {
	Uid      int    `json:"uid"`
	Key      string `json:"key"` // passwd:reset:{uuid} = uid
	Password string `json:"password,omitempty"`
}

type PasswordResetResp struct {
	app.Response
	Uid      int    `json:"uid"`
	Password string `json:"password,omitempty"`
}

// ChangePassword
/*

修改密码
need:old & new
*/
func ChangePassword(ctx iris.Context) {
	log.Debug("change password")
	var req PasswordReq
	var resp app.Response
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("req err")
		log.Info(err)
		resp.Code = -1
		resp.Message = "req err"
		common.ResponseJSON(ctx, resp)
		return
	}
	//check
	if req.Uid == 0 {
		log.Info("uid nil")
		resp.Code = -1
		resp.Message = "uid nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Old) <= 0 {
		log.Info("Old Password nil")
		resp.Code = -1
		resp.Message = "Old Password nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.New) <= 0 {
		log.Info("New Password nil")
		resp.Code = -1
		resp.Message = "New Password nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	// user
	u, err := userService.GetUser(req.Uid)
	if err != nil {
		log.Info("unknown uid:%d", req.Uid)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	// TODO verify old
	success := password.Compare(u, req.Old)
	if !success {
		log.Info("Old Password err")
		log.Info(success)
		resp.Code = -1
		resp.Message = "Old Password err"
		common.ResponseJSON(ctx, resp)
		return
	}
	// TODO change passwd
	newPasswd, err := password.Generate(u, req.New)
	if err != nil {
		log.Info("gen password failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	err = userService.ChangePassword(u, newPasswd)
	if err != nil {
		log.Info("update password failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

// ChangePasswordWithKey
/*

使用授权key修改密码
need:key & new

key: means that program has auth to change passwd
*/
func ChangePasswordWithKey(ctx iris.Context) {

	// TODO verify key
	// TODO del key
	// TODO change passwd
}

// ResetPassword
/*

使用授权key重置密码

*/
func ResetPassword(ctx iris.Context) {

	var err error
	var req PasswordResetReq
	var resp PasswordResetResp
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read param fail")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}
	// TODO verify key
	var forUid int
	err = app.GetCache().Get(req.Key, &forUid)
	if err != nil {
		log.Info("unknown key")
		log.Info(err)
		resp.Code = respcode.EditPasswordFailed
		common.ResponseJSON(ctx, resp)
		return
	}
	//if len(forUid) == 0 {
	//	log.Info("unknown key")
	//	log.Info(err)
	//	resp.Code = respcode.EditPasswordFailed
	//	common.ResponseJSON(ctx, resp)
	//	return
	//}

	// TODO del key

	// TODO gen random passwd
	u, err := userService.GetUser(forUid)
	if err != nil {
		log.Info("unknown uid:%d", forUid)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	var plainPassword = common.RandTxt(12)
	newPasswd, err := password.Generate(u, plainPassword)
	if err != nil {
		log.Info("gen password failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}
	// TODO change passwd
	err = userService.ChangePassword(u, newPasswd)
	if err != nil {
		log.Info("update password failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	resp.Password = plainPassword
	common.ResponseJSON(ctx, resp)
}

// ResetPwd
/*

管理员重置密码

*/
func ResetPwd(ctx iris.Context) {

	var err error
	var req PasswordResetReq
	var resp PasswordResetResp
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read param fail")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	// gen random passwd
	u, err := userService.GetUser(req.Uid)
	if err != nil {
		log.Info("unknown uid:%d", req.Uid)
		log.Info(err)
		resp.Code = common.RespCodeNotFound
		common.ResponseJSON(ctx, resp)
		return
	}
	var plainPassword = common.RandTxt(12)
	newPasswd, err := password.Generate(u, plainPassword)
	if err != nil {
		log.Info("gen password failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}
	// change passwd
	err = userService.ChangePassword(u, newPasswd)
	if err != nil {
		log.Info("update password failed")
		log.Info(err)
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = common.RespCodeSuccess
	resp.Password = plainPassword
	common.ResponseJSON(ctx, resp)
}