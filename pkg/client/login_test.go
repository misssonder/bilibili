package client

import (
	"testing"
)

func TestLogin(t *testing.T) {
	client := &Client{}
	statuses, err := client.LoginWithQrCode()
	if err != nil {
		t.Error(err)
		return
	}
	for status := range statuses {
		println(status)
	}

}
