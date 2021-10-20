package messager

import "github.com/lishimeng/go-log"

func SendMail(message MessageSdk, sender, tpl, subject string, params interface{}, receiverMail string) (response Response, err error) {
	var request Request
	request.Params = params
	request.Subject = subject
	request.Sender = sender
	request.Receiver = receiverMail
	request.Template = tpl

	response, err = message.Send(request)
	if err != nil {
		log.Info(err)
		return
	}
	return
}
