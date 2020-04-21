package setup

import (
	"context"
	"github.com/lishimeng/auth/internal/cache"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/auth/internal/token"
	"time"
)

func Setup(ctx context.Context) (err error) {

	modules := []func(context.Context) error{setupToken, setupRedisCache}

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
		time.Hour*time.Duration(etc.Config.Token.Expire)))
	return
}

func setupRedisCache(ctx context.Context) (err error) {
	cache.Init(
		ctx,
		etc.Config.Redis.Addr,
		etc.Config.Redis.Password,
		etc.Config.Redis.Db,
		)
	return
}
