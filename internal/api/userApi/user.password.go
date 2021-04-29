package userApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/go-log"
)

type PasswordReq struct {
	Uid int `json:"uid"`
	Old string `json:"old"`
	New string `json:"new"`
}

type PasswordResetReq struct {
	Uid int `json:"uid"`
	Key string `json:"key"` // passwd:reset:{uuid} = uid
	Password string `json:"password,omitempty"`
}

type PasswordResetResp struct {
	app.Response
	Uid int `json:"uid"`
	Password string `json:"password,omitempty"`
}

// ChangePassword
/*

修改密码
need:old & new
 */
func ChangePassword(ctx iris.Context) {

	// TODO verify old
	// TODO change passwd
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
	var forUid string
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
	// TODO change passwd
}
