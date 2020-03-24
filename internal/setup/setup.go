package setup

import (
	"context"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/token"
	"github.com/lishimeng/go-libs/jwt"
	"time"
)

func Setup(ctx context.Context) (err error) {

	modules := []func(context.Context)error{ setupToken}

	for _, m := range modules {
		err = m(ctx)
		if err != nil {
			break
		}
	}
	return
}

func setupToken(_ context.Context) (err error) {
	token.Init(jwt.New([]byte(etc.Config.Token.Secret),
		etc.Config.Token.Issuer,
		time.Hour * time.Duration(etc.Config.Token.Expire)))
	return
}
