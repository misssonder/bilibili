package client

import (
	"os"
	"testing"

	"github.com/misssonder/bilibili/internal/util"
)

func TestClient_GetVideoInfo(t *testing.T) {
	client := &Client{}
	info, err := client.GetVideoInfo("BV117411r7R1")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshalIndent(info))
}

func TestClient_PlayUrl(t *testing.T) {
	client := &Client{}
	responses, err := client.LoginWithQrCode(os.Stdout)
	if err != nil {
		t.Error(err)
		return
	}
	for resp := range responses {
		switch resp.LoginStatus {
		case LoginSuccess:
			client.SetCookie(resp.Cookie)
			break
		case LoginExpired:
			t.Log(LoginExpired)
			return
		default:
			continue
		}
	}
	resp, err := client.PlayUrl("https://www.bilibili.com/video/BV1XG4y1X7HV/?spm_id_from=333.1007.tianma.1-3-3.click&vd_source=76326787bdfce30577382b0e7e18f35c", "973082699", 120)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshalIndent(resp))
}
