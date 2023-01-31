package client

import (
	"testing"

	"github.com/misssonder/bilibili/internal/util"
)

func TestLogin(t *testing.T) {
	client := &Client{}
	qrcode, err := client.GenerateQrcode()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(qrcode))

	pollQrcode, err := client.PollQrcode(qrcode.Data.QrcodeKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(pollQrcode))
}
