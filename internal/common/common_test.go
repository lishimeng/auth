package common

import "testing"

func TestRandTxt(t *testing.T) {

	var txt string
	txt = RandTxt(12)
	t.Logf("txt is:->[%s]<-", txt)
}
