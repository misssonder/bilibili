package client

import (
	"os"
	"testing"
	
	"github.com/misssonder/bilibili/internal/util"
)

func TestLogin(t *testing.T) {
	client := &Client{}
	statuses, err := client.LoginWithQrCode(os.Stdout)
	if err != nil {
		t.Error(err)
		return
	}
	for status := range statuses {
		println(status)
	}

}

func TestClient_NavInfo(t *testing.T) {
	client := &Client{}
	statuses, err := client.LoginWithQrCode(os.Stdout)
	if err != nil {
		t.Error(err)
		return
	}
	for range statuses {
	}

	info, err := client.NavInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(info))
}
