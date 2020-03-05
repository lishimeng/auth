package setup

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/token"
	"github.com/lishimeng/auth/internal/web/api"
	"github.com/lishimeng/go-libs/jwt"
	"github.com/lishimeng/go-libs/log"
	server "github.com/lishimeng/go-libs/web"
	"time"
)

func Setup() (err error) {

	modules := []func()error{ setupToken, setupWeb}

	for _, m := range modules {
		err = m()
		if err != nil {
			break
		}
	}
	return
}

func setupToken() (err error) {
	token.Init(jwt.New([]byte(etc.Config.Token.Secret),
		etc.Config.Token.Issuer,
		time.Hour * time.Duration(etc.Config.Token.Expire)))
	return
}

func setupWeb() (err error) {

	go func() {
		log.Info("start web server")
		s := server.New(server.ServerConfig{
			Listen: etc.Config.Web.Listen,
		})
		s.OnErrorCode(404, func(ctx iris.Context) {

			_, _ = ctx.Text("not found(404~)")
		})
		s.RegisterComponents(api.TokenApi)
		err = s.Start()
		if err != nil {
			log.Info(err)
		}
		log.Info("stop web server")
	}()
	return
}