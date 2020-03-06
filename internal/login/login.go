package login

type UserInfo struct {
	Uid string
	LoginName string
}

func Login(name, password string) (ui UserInfo, err error) {

	ui = UserInfo{
		Uid:       name,
		LoginName: name,
	}
	return
}