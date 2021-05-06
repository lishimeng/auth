package common

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris/v12"
)

const (
	RespCodeSuccess  = 200
	RespCodeNotFound = 404

	RespCodeInternalError = 500
)

const (
	RespMsgNotFount = "not found"
	RespMsgIdNum    = "id must be a int value"
)

func ResponseJSON(ctx iris.Context, j interface{}) {
	_, _ = ctx.JSON(j)
}

const (
	DefaultTimeFormatter = "2006-01-02:15:04:05"
	DateFormatter        = "2006-01-02"
	DefaultCodeLen       = 16
)

func FormatTime(t time.Time) (s string) {
	s = t.Format(DefaultTimeFormatter)
	return
}
func GetRandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

const (
	UserStatusActivate   = 20
	UserStatusDeActivate = 10
)
