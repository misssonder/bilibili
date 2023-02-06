package client

import (
	"testing"

	"github.com/misssonder/bilibili/internal/util"
)

func TestClient_SeasonSection(t *testing.T) {
	client := &Client{}
	resp, err := client.SeasonSection("", "729217")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(util.MustMarshalIndent(resp))
}
