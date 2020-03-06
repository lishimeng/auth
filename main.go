package main

import (
	"fmt"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/setup"
	"github.com/lishimeng/go-libs/log"
	"github.com/lishimeng/go-libs/shutdown"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ctx := shutdown.Context()

	etc.SetEnvPath([]string{".", "/etc/auth"})
	etc.SetConfigName("config.toml")
	err := etc.LoadEnvs()
	if err != nil {
		log.Info(err)
		return
	}

	err = setup.Setup(ctx)
	if err != nil {
		log.Info(err)
		return
	}

	shutdown.WaitExit(&shutdown.Configuration{
		BeforeExit: func(s string) {
			log.Info(s)
		},
	})
}
