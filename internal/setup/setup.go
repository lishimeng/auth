package setup

import (
	"context"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/jwt"
	"github.com/lishimeng/auth/internal/password"
	"github.com/lishimeng/auth/internal/token"
	"time"
)

func Setup(ctx context.Context) (err error) {

	modules := []func(context.Context) error{setupToken, setupRefreshKey}

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

func setupRefreshKey(ctx context.Context) (err error) {

	go func() {
		var d = time.Minute * 21
		var t = time.NewTimer(d)
		defer func() {
			t.Stop()
		}()
		password.RefreshPasswordKey()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				password.RefreshPasswordKey()
				t.Reset(d)
			}
		}
	}()
	return err
}
