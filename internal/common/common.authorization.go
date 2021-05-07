package common

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/auth/internal/token"
	"strings"
)

const (
	AuthHeaderKey = "Authorization"
	AuthBearerKey = "Bearer"

	AuthAccessKey = "accessToken"
)

func Authorization(ctx iris.Context) (claims *jwt.Claims, success bool) {
	str := GetAuthorization(ctx)
	claims, success = VerifyToken(str)
	return
}

func GetAuthorization(ctx iris.Context) (tokenStr string) {
	// Authorization Bearer
	bearer := ctx.GetHeader(AuthHeaderKey)
	if len(bearer) > 0 {
		if strings.HasPrefix(bearer, AuthBearerKey) {
			tokenStr = strings.Replace(bearer, AuthBearerKey, ContentBlank, 1)
		}
	}

	// accessToken:
	// schema://domain/path?accessToken=
	if len(tokenStr) == 0 {
		tokenStr = ctx.URLParamTrim(AuthAccessKey)
	}
	return
}

func VerifyToken(tokenStr string) (claims *jwt.Claims, success bool) {
	claims, success = token.Verify(tokenStr)
	return
}
