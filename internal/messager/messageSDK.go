package messager

type MessageSdk interface {
	send(request Request)(code interface{}, err error)
	enable(debug bool)
	getDebug()(debug bool)
}
