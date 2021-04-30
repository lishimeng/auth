package common

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"math/rand"
	"time"
)

const (
	CacheRsaBits   = 1024
	CacheSignInKey = "global:sign_in:key"
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
	DateFormatterNoSplit = "20060102"
	DefaultCodeLen       = 16
)

func FormatTime(t time.Time) (s string) {
	s = t.Format(DefaultTimeFormatter)
	return
}
func RandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()")

func RandTxt(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62+10)]
	}
	return string(b)
}
