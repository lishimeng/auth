package messager

type SimpleSender struct {

}

func (s SimpleSender) SendMail(messager MessageSdk, sender, tpl, subject string, params interface{}, receiverMail string) {
	var request Request
	request.Params = params
	request.Subject = subject
	request.Sender = sender
	request.Receiver = receiverMail
	request.Template = tpl

	_, err := messager.send(request)
	if err != nil {
		return
	}
}
