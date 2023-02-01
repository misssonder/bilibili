package client

import (
	"os"
	"testing"
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
