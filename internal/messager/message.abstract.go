package messager

type Message struct {
	path string
	host string
	debug bool
}


func (m *Message) Send(request Request)(code interface{}, err error) {
	var url = m.getReqUrl()
	code = 0

	code, err = m.SendMail(url,request)
	if err != nil {
		return -1, err
	}
	return
}

func (m *Message)getReqUrl()(url string) {
	return m.getHost() + m.path
}

func (m *Message)getHost()(host string) {
	return m.host
}

func (m *Message)setPath(path string){
	m.path = path
}
func (m *Message)setHost(host string){
	m.path = host
}
func (m *Message)enable(debug bool){
	m.debug = debug
}
func (m *Message)getDebug()(debug bool){
	return m.debug
}
