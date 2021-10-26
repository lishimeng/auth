package messager

type Message struct {
	path  string
	host  string
	Debug bool
}

func NewMessage(host, path string, debug bool) (m MessageSdk) {
	msg := Message{
		path: path,
		host: host,
		Debug: debug,
	}
	m = &msg
	return
}

func (m Message) Send(request Request) (response Response, err error) {
	var url = m.getReqUrl()
	response, err = m.SendMail(url, request)
	if err != nil {
		return
	}
	return
}

func (m Message) getReqUrl() (url string) {
	return m.host + m.path
}

func (m Message) setPath(path string) {
	m.path = path
}
func (m Message) setHost(host string) {
	m.host = host
}
