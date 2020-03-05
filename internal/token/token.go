package token

import (
	"github.com/lishimeng/go-libs/jwt"
)

type Token struct {
	jwt.Claims
	Jwt string `json:"jwt"`
}

var jwtHandler *jwt.Handler

func Init(handler jwt.Handler) {
	jwtHandler = &handler
}

func Gen(uid string, loginType int32, ) (token Token, success bool) {

	req := jwt.Token{
		BaseToken: jwt.BaseToken{UID: uid, LoginType: loginType},
		Audience:  uid,
	}
	var claims *jwt.Claims
	var signedToken string
	claims, signedToken, success = jwtHandler.GenToken(req)
	if success {
		token = Token{
			Claims: *claims,
			Jwt:    signedToken,
		}
	}
	return
}

func Verify(jwtStr string, forUid string) (success bool) {
	c, success := jwtHandler.VerifyToken(jwtStr)
	if success {
		if len(forUid) > 0 {
			success = forUid == c.UID
		}
	}
	return
}