package client

import (
	"github.com/misssonder/bilibili/internal/util"
	"github.com/misssonder/bilibili/pkg/qrcode"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	client := &Client{}
	qrcodeResp, err := client.GenerateQrcode()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(qrcodeResp))

	if err = qrcode.Generate(qrcodeResp.Data.URL, qrcode.Low, os.Stdout); err != nil {
		return
	}
	for {
		pollQrcode, err := client.PollQrcode(qrcodeResp.Data.QrcodeKey)
		if err != nil {
			t.Error(err)
			return
		}
		if pollQrcode.Data.Code != 0 {
			continue
		} else {
			t.Log(util.MustMarshal(pollQrcode))
			break
		}
	}

}
