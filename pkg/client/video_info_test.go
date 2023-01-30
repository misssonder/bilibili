package client

import (
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
