package main

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/cache"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/auth/cmd"
	"github.com/lishimeng/auth/internal/api"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/etc"
	"github.com/lishimeng/auth/internal/setup"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"time"
)
import _ "github.com/lib/pq"

func main() {
	orm.Debug = true

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(cmd.AppName)
	fmt.Println(cmd.Version)

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		redisOpts := cache.RedisOptions{
			Addr: etc.Config.Redis.Addr,
			Password: etc.Config.Redis.Password,
		}
		cacheOpts := cache.Options{
			MaxSize: 10000,
			Ttl:     time.Hour*24,
		}
		builder.EnableDatabase(dbConfig.Build(),
			new(model.AuthUser),
			new(model.AuthUserOrganization),
			new(model.AuthRoleOrganization),
			new(model.AuthOrganization),
			new(model.AuthUserRoles),
			new(model.AuthRole)).
			SetWebLogLevel("debug").
			EnableCache(redisOpts, cacheOpts).
			EnableWeb(etc.Config.Web.Listen, api.Route).ComponentAfter(setup.Setup)
		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
