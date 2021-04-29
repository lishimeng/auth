package userApi

import "github.com/kataras/iris/v12"

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

	// TODO verify key
	// TODO del key
	// TODO gen random passwd
	// TODO change passwd
}
