package password

import (
	"github.com/lishimeng/auth/internal/db/model"
	"testing"
	"time"
)

var now = time.Now()

var u = model.AuthUser{
	Pk:              model.Pk{Id: 1},
	UserNo:          "hahaha",
	UserName:        "",
	Email:           "",
	Phone:           "",
	Password:        "",
	TableChangeInfo: model.TableChangeInfo{},
}

var password = "123456"
var passwordWrong = "123456xxx"

func TestGenerate(t *testing.T) {
	passwdEnc, err := Generate(u, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("password: %s", passwdEnc)
}

func TestCompare001(t *testing.T) {
	passwdEnc, err := Generate(u, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("password: %s", passwdEnc)
	u.Password = passwdEnc

	success := Compare(u, password)
	if !success {
		t.Fatal("password is not match")
	}
}

func TestCompare002(t *testing.T) {
	passwdEnc, err := Generate(u, password)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("password: %s", passwdEnc)
	u.Password = passwdEnc

	success := Compare(u, passwordWrong)
	if success {
		t.Fatal("assert password is not match")
	}
}
