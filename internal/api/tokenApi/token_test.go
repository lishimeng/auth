package tokenApi

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/app-starter"
	"strings"
	"testing"
)

func TestResponse(t *testing.T) {
	resp := app.Response{
		Code:    0,
		Message: "",
	}
	bs, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(string(bs))
}

func TestBearer(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJVMTM0NDJzc3MiLCJsb2dpbl90eXBlIjoyLCJhdWQiOiJVMTM0NDJzc3MiLCJleHAiOjE1ODM0MjM4MDAsImlzcyI6InRjcDovL2xvY2FsaG9zdDoxODgzIn0.ss4-wDz_iLWxL3MaCWTwQymTddmJJxCW7vcUJsXrbVE"
	bearer := "Bearer " + tokenStr

	str := strings.Replace(bearer, "Bearer ", "", 1)

	if tokenStr != str {
		t.Fatal(str)
		return
	}
}
