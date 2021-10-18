package sendMailApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/messager"
	"github.com/lishimeng/auth/internal/respcode"
	"github.com/lishimeng/go-log"
)

type SenderReq struct {
	Mail string `json:"mail,omitempty"`
}
type MailParams struct {
	Code string `json:"code,omitempty"`
}
func Send(ctx iris.Context) {
	var err error
	var req SenderReq
	var resp app.Response

	// 验证邮箱
	err = ctx.ReadJSON(&req)
	if err != nil || !(len(req.Mail) > 0) {
		log.Info("req err")
		resp.Code = respcode.SendMailFailed
		resp.Message = "Wrong request parameters."
		common.ResponseJSON(ctx, resp)
		return
	}
	// 生成随机字符串-验证码（6位？）
	var code = common.RandCode(6)
	log.Info("Captcha: %s", code)
	// 将验证码放入缓存
	err = app.GetCache().Set(req.Mail, code)
	if err != nil {
		log.Info("Cache err")
		resp.Code = respcode.SendMailFailed
		resp.Message = "Cache err"
		common.ResponseJSON(ctx, resp)
		return
	}

	var params MailParams
	params.Code = code
	// 调用 message 接口发送邮件验证码
	var simpleSender messager.SimpleSender
	var message messager.MessageSdk
	simpleSender.SendMail(
		message,
		etc.Config.Mail.Sender,
		"243w54e65r76t879y8",
		"Captcha | CM Venture Capital Proprietary Database",
		params,
		req.Mail)
	// 返回结果-发送验证码成功
	resp.Code = 200
	common.ResponseJSON(ctx, resp)
}
