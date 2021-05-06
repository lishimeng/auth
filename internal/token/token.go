package token

import (
	"github.com/lishimeng/auth/internal/jwt"
	"time"
)

type Token struct {
	jwt.Claims
	Jwt string `json:"jwt"`
	Expire time.Duration `json:"-"`
}

var jwtHandler *jwt.Handler

func Init(handler jwt.Handler) {
	jwtHandler = &handler
}

func Gen(uid, oid int, loginType int32, expire time.Duration) (token Token, success bool) {

	var req jwt.TokenReq
	req.UID = uid
	req.OID = oid
	req.Type = loginType
	if expire > 0 {
		req.Expire = expire
	} else {

	}

	var claims *jwt.Claims
	var tokenExpire time.Duration
	var signedToken string
	claims, tokenExpire, signedToken, success = jwtHandler.GenToken(req)
	if success {
		token = Token{
			Claims: *claims,
			Expire: tokenExpire,
			Jwt:    signedToken,
		}
	}
	return
}

func Verify(jwtStr string, forUid int) (success bool) {
	c, success := jwtHandler.VerifyToken(jwtStr)
	if success {
		success = forUid == c.UID
	}
	return
}
