package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {

	h := New([]byte("secret"), "AAA.com", time.Hour*3*24)
	c, _, token, success := h.GenToken(TokenReq{
		BaseToken: BaseToken{
			UID:  1234,
			OID:  1,
			Type: 1,
		},
		//Audience:  "U13413",
		//Subject:   "AAA.com/user1",
	})
	fmt.Println(token)
	fmt.Println(success)
	if len(token) <= 0 {
		t.Fatal()
		return
	}
	if !success {
		t.Fatal()
		return
	}

	c, success = h.VerifyToken(token)
	if !success {
		t.Fatal("verify failed")
		return
	}

	bs, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(bs))
}
