package qrcode

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	err := Generate("https://passport.bilibili.com/h5-app/passport/login/scan?navhide=1\\u0026qrcode_key=d8d8decc486a9aa64b72d15cf3a6c307\\u0026from=", Low, os.Stdout)
	if err != nil {
		t.Error(err)
		return
	}
}
