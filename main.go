package main

import (
	"fmt"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/setup"
	"github.com/lishimeng/auth/internal/web/api"
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-app-shutdown"
	"github.com/lishimeng/go-log"
)

func main() {

	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ctx := shutdown.Context()
	application := app.New(ctx)

	err = application.LoadConfig(&etc.Config, "config.toml", ".", "/etc/auth")
	if err != nil {
		log.Info(err)
		return
	}

	err = setup.Setup(ctx)
	if err != nil {
		log.Info(err)
		return
	}

	err = setup.Setup(ctx)
	if err != nil {
		log.Info(err)
		return
	}

	err = application.
		EnableWeb(etc.Config.Web.Listen, api.TokenApi).
		Start()
	if err != nil {
		log.Info(err)
		return
	}

	shutdown.WaitExit(&shutdown.Configuration{
		BeforeExit: func(s string) {
			log.Info(s)
		},
	})
	time.Sleep(time.Second)
}
