package client

import (
	"os"
	"testing"

	"github.com/misssonder/bilibili/internal/util"
)

func TestLogin(t *testing.T) {
	client := &Client{}
	resps, err := client.LoginWithQrCode(os.Stdout)
	if err != nil {
		t.Error(err)
		return
	}
	for resp := range resps {
		if resp.LoginStatus == LoginSuccess {
			t.Log(util.MustMarshal(resp.Cookie))
		}
	}

}

func TestClient_NavInfo(t *testing.T) {
	client := &Client{}
	resps, err := client.LoginWithQrCode(os.Stdout)
	if err != nil {
		t.Error(err)
		return
	}
	for resp := range resps {
		if resp.LoginStatus == LoginSuccess {
			client.SetCookie(resp.Cookie)
		}
	}
	info, err := client.NavInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(info))
}
