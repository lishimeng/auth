package token

import (
	"github.com/lishimeng/auth/internal/jwt"
	"testing"
	"time"
)

func TestVerify(t *testing.T) {
	var secret = "2lP2ybr5c3ZKjH2SlZk7krbrv8T9eFxQ"
	var expire = 24
	var issuer = "thingple.com"
	var testStr = ""
	Init(jwt.New([]byte(secret), issuer, time.Hour*time.Duration(expire)))
	v, suc := Gen(1,2,3,time.Duration(0))
	if !suc {
		t.Fatal("gen token failed")
	}
	testStr = v.Jwt
	t.Logf("jwt is %s", testStr)
	v2, suc := Verify(testStr)
	if !suc {
		t.Fatal("verify token failed")
	}
	if v2.UID != 1 {
		t.Fatal("uid not match")
	}
	if v2.OID != 2 {
		t.Fatal("oid not match")
	}
}
