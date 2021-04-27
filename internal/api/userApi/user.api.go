package userApi

import "github.com/kataras/iris/v12"

type Req struct {
	LoginName string `json:"loginName,omitempty"`
	Password string `json:"password,omitempty"`
}

func SignIn(ctx iris.Context) {

	// TODO 校验用户
	// TODO 创建jwt
	// TODO 返回jwt, role_list
	// TODO role_list 存入cache
	// TODO jwt 存入cache
}

func Logout(ctx iris.Context) {
	ctx.GetHeader("")
}