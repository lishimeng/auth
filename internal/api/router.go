package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/lishimeng/auth/internal/api/authRoleApi"
	"github.com/lishimeng/auth/internal/api/authUserApi"
	"github.com/lishimeng/auth/internal/api/registerUserApi"
	"github.com/lishimeng/auth/internal/api/sendMailApi"
	"github.com/lishimeng/auth/internal/api/tokenApi"
	"github.com/lishimeng/auth/internal/api/userApi"
)

func Route(app *iris.Application) {
	app.Use(recover.New())
	route(app.Party("/api"))
}

func route(root iris.Party) {
	token(root.Party("/token"))
	user(root.Party("/user"))
	authRoles(root.Party("/authRoles"))
	authUser(root.Party("/authUser"))
	registerUser(root.Party("/registerUserApi"))
	send(root.Party("/mailCode"))
}

func registerUser(p iris.Party) {
	p.Post("/", registerUserApi.Register) // 用户注册
}

func send(p iris.Party) {
	p.Post("/", sendMailApi.Send) // 发送邮箱验证码
}

func token(p iris.Party) {
	p.Post("/verify", tokenApi.VerifyToken)
}

func user(p iris.Party) {
	p.Post("/sign_in", userApi.SignIn)               // 登录
	p.Post("/sign_in_card", userApi.SignInCard)      // 智能卡登录
	p.Post("/logout", Authorization, userApi.Logout) // 退出
	p.Get("/info/{id}", userApi.GenUserInfo)         // 用户信息

	p.Post("/password/change", userApi.ChangePassword)
	p.Post("/password/change_with_key", userApi.ChangePasswordWithKey)
	//p.Post("/password/reset", userApi.ResetPassword)
	p.Post("/password/reset", userApi.ResetPwd)
}

// authUser
func authUser(p iris.Party) {
	p.Post("/add", Authorization, authUserApi.Add)                          // 添加用户
	p.Get("/", Authorization, authUserApi.GetUserList)                      // 用户列表
	p.Get("/{id}", Authorization, authUserApi.GetUserInfo)                  // 用户信息id
	p.Put("/{id}", Authorization, authUserApi.UpdateUserInfo)               // 更新用户
	p.Put("/{id}/status", Authorization, authUserApi.UpdateUserStatus)      // 改用户状态
	p.Put("/roles/change/{id}", Authorization, authUserApi.UpdateUserRoles) // deprecate
}

// authRoles
func authRoles(p iris.Party) {
	p.Get("/", Authorization, authRoleApi.GetRoleList)                              // 角色列表
	p.Post("/", Authorization, authRoleApi.Add)                      // 添加角色
	p.Delete("/{id}", Authorization, authRoleApi.Del)                // 删除角色(id:角色关系表id)
	p.Delete("/{uid}/{rid}", Authorization, authRoleApi.DelUserRole) // 删除用户的角色,需要通过user和role查询
}
