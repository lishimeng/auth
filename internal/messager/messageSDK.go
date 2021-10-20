package messager

type MessageSdk interface {
	Send(request Request)(response Response, err error)
}
