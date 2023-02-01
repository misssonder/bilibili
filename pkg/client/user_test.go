package client

import (
	"testing"

	"github.com/misssonder/bilibili/internal/util"
)

func TestClient_MySpaceInfo(t *testing.T) {
	client := &Client{}
	statuses, err := client.LoginWithQrCode()
	if err != nil {
		return
	}
	for range statuses {
		continue
	}
	info, err := client.MySpaceInfo()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshal(info))
}
