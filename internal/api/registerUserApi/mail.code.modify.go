package registerUserApi

import (
	"github.com/lishimeng/app-starter"
)

func ModifyMailCode(email,code string) (isCode bool, err error) {
	var cacheCode string
	isCode = false

	err = app.GetCache().Get(email,&cacheCode)
	if err != nil {
		return
	}

	if code != cacheCode {
		isCode = false
	}

	isCode = true
	return
}
