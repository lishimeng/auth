package messager

import (
	"bytes"
	"encoding/json"
	"github.com/lishimeng/go-log"
	"io/ioutil"
	"net/http"
	"time"
)

func (m *Message) SendMail(url string, request Request) (response Response, err error) {
	log.Info("sendMail url: %s", url)
	log.Info("sendMail to: %s", request.Receiver)
	if m.Debug {
		log.Info("Debug sendMail")
		return
	}

	client := &http.Client{Timeout: 8 * time.Second}
	jsonStr, _ := json.Marshal(request)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(result, &response)
	if err != nil {
		return
	}
	log.Info("sendMail response: %v", response)
	return

}
