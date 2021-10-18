package messager

import (
	"bytes"
	"encoding/json"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/go-log"
	"io/ioutil"
	"net/http"
	"time"
)

var message Message
const path = "/api/send/mail"

func (Message) SendMail(url string, request Request)(code interface{}, err error){
	message.setPath(path)
	message.setHost(etc.Config.Mail.Host)
	message.enable(etc.Config.Mail.Debug)

	log.Info("sendMail url: %s" + url)
	log.Info("sendMail to: %s" + request.Receiver)
	if message.getDebug() {
		log.Info("Debug sendMail")
		return
	}

	client := &http.Client{Timeout: 8 * time.Second}
	jsonStr, _ := json.Marshal(request)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	log.Info("sendMail response: %s", result)
	return

}
