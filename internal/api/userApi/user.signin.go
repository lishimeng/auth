package userApi

import (
	"encoding/base64"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	myrsa "github.com/lishimeng/auth/internal/api/rsaApi"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/service/roleService"
	"github.com/lishimeng/auth/internal/db/service/userService"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/auth/internal/token"
	"github.com/lishimeng/auth/internal/utils"
	"github.com/lishimeng/go-log"
)

func SignIn(ctx iris.Context) {

	var req Req
	var resp Resp
	var err error
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = common.RespCodeInternalError
		common.ResponseJSON(ctx, resp)
		return
	}
	// TODO 前端RSA-pubkey加密，后端RSA-prikey解密
	// plainPwd := handleFrontEncodingPwdToPlainText()

	// 校验用户
	u, err := userService.SignIn(req.LoginName, req.Password)
	if err != nil {
		log.Info("sign in failed")
		log.Info(err)
		resp.Code = respcode.SignInFailed
		common.ResponseJSON(ctx, resp)
		return
	}
	auo, err := userService.GetAuthUserOrg(u)
	if err != nil {
		log.Info("no auth org info")
		log.Info(err)
	} else {
		resp.OrgId = auo.OrgId
	}
	// org -> cache
	// 创建jwt
	t, success := token.Gen(u.Id, auo.OrgId, 1, 0)
	if !success {
		log.Info("create jwt failed")
		resp.Code = respcode.SignInFailed
		common.ResponseJSON(ctx, resp)
		return
	}

	// role_list
	aurs, err := roleService.GetUserRoles(u.Id)
	if err != nil {
		log.Info("get auth role fail uid:%d", u.Id)
		log.Info(err)
	} else {
		for _, r := range aurs {
			resp.Roles = append(resp.Roles, r.RoleId)
		}
	}
	// TODO role_list 存入cache
	// jwt 存入cache
	err = app.GetCache().Set(t.Jwt, t.UID)
	if err != nil {
		log.Info(err)
	}
	// TODO 返回jwt

	resp.Code = common.RespCodeSuccess
	resp.Jwt = t.Jwt
	resp.Uid = u.Id
	common.ResponseJSON(ctx, resp)

}

// 处理前端加密文解析为明文
func handleFrontEncodingPwdToPlainText(encodingPwd string) (pwd string) {
	// 从redis中获取密钥进行解密
	priKey := ""
	err := app.GetCache().Get(myrsa.CONANT_CACHE_KEY, &priKey)
	if err != nil {
		log.Debug("get prikey failed %s", err.Error())
		return
	}

	// TODO 前端使用公钥加密的密文这里要进行一次base64解码一下！！！
	b, err2 := base64.StdEncoding.DecodeString(encodingPwd)
	if err2 != nil {
		log.Debug("base64.StdEncoding error %s", err2.Error())
		return
	}

	// RSA解密
	decodeStr, err3 := utils.DecodingByPrivateKey(priKey, b)
	if err3 != nil {
		log.Debug("DecodingByPrivateKey error %s", err3)
	}

	// 返回明文
	return string(decodeStr)
}
