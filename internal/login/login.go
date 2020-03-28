package login

import "errors"

type UserInfo struct {
	Uid string
	LoginName string
}

var systemUser = map[string]string{}

func InitSysUser(name, password string) {
	systemUser[name] = password
}

func Login(name, password string) (ui UserInfo, err error) {

	ui = UserInfo{}
	v, ok := systemUser[name]
	if ok && v==password {
		ui.Uid = name
		ui.LoginName = name
		return
	}
	err = errors.New("invalid" + name)
	return
}