package messager

type Request struct {
	Subject  string      `json:"subject,omitempty"`
	Sender   string      `json:"sender,omitempty"`
	Receiver string      `json:"receiver,omitempty"`
	Template string      `json:"template,omitempty"`
	Params   interface{} `json:"params,omitempty"`
}
