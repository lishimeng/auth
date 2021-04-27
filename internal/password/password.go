package password

import (
	"fmt"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"golang.org/x/crypto/bcrypt"
)

func Generate(u model.AuthUser, plaintext string) (p string, err error) {

	bs, err := bcrypt.GenerateFromPassword([]byte(genPlainPassword(u, plaintext)), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	p = string(bs)
	return
}

func Compare(u model.AuthUser, plaintext string) (success bool) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(genPlainPassword(u, plaintext)))
	success = err == nil
	return
}

func genPlainPassword(u model.AuthUser, plaintext string) (s string) {
	ts := u.CreateTime.Format(common.DefaultTimeFormatter)
	return fmt.Sprintf("%d.%s.%s_.%s", u.Id, u.Email, ts, plaintext)
}