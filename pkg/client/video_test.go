package client

import (
	"os"
	"testing"

	"github.com/misssonder/bilibili/internal/util"
	"github.com/misssonder/bilibili/pkg/video"
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
	bvID := "BV117411r7R1"
	info, err := client.GetVideoInfo(bvID)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := client.PlayUrl(bvID, int64(info.Data.Pages[0].Cid), Qn4k, FnvalHDR|Fnval4K)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshalIndent(resp))
}

func TestClient_PlayUrlV2(t *testing.T) {
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
	id := "https://www.bilibili.com/bangumi/play/ep759371"
	epID, err := video.ExtractEpID(id)
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := client.PlayUrlV2(epID, Qn4k, FnvalHDR|Fnval4K)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshalIndent(resp))
}
