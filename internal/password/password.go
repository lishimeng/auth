package password

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"golang.org/x/crypto/bcrypt"
	"hash"
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
	salt := fmt.Sprintf("%d.%s.%s_.%s", u.Id, u.Email, ts, plaintext)
	return digest(plaintext, salt, 10)
}

func digest(plaintext, salt string, loop int) (dig string) {
	dig = _digest(sha512.New(), []byte(plaintext), []byte(salt), loop)
	return
}
func _digest(dig hash.Hash, plaintext, salt []byte, loop int) (txt string) {
	var tmp = plaintext
	for i := 0; i < loop; i++ {
		dig.Write(tmp)
		dig.Write(salt)
		tmp = dig.Sum(nil)
	}
	return base64.StdEncoding.EncodeToString(tmp)
}