package userService

import (
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/db/repo"
	"github.com/lishimeng/auth/internal/password"
)

func SignIn(name, plaintext string) (u model.AuthUser, err error) {
	var ctx = app.GetOrm()
	u, err = repo.GetUserAsLogin(*ctx, name)
	if err != nil {
		return
	}

	success := password.Compare(u, plaintext)
	if !success {
		err = fmt.Errorf("password wrong")
	}
	return
}

