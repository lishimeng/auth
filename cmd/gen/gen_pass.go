package main

import (
	"github.com/lishimeng/auth/internal/common"
	"github.com/lishimeng/auth/internal/db/model"
	"github.com/lishimeng/auth/internal/password"
	"log"
	"time"
)

func main() {
	ct, _ := time.Parse(common.DateFormatter, "2021-07-11")
	pwd := "123456"
	u := model.AuthUser{
		Pk:              model.Pk{Id: 1},
		UserNo:          "00001",
		UserName:        "00001",
		Email:           "admin@thingple.com",
		Phone:           "12345678901",
		Password:        pwd,
		TableChangeInfo: model.TableChangeInfo{
			Status:     1,
			CreateTime: ct,
			UpdateTime: ct,
		},
	}
	passCode, _ := password.Generate(u, pwd)
	log.Printf("明文:%s", pwd)
	log.Printf("密文:%s", passCode)
}
