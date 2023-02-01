package client

import (
	"encoding/json"
	"fmt"
	"github.com/misssonder/bilibili/pkg/errors"
	"github.com/misssonder/bilibili/pkg/qrcode"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	generateQrCodeUrl = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	pollQrCodeUrl     = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
)

type GenerateQrCodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		URL       string `json:"url"`
		QrcodeKey string `json:"qrcode_key"`
	} `json:"data"`
}

type PollQrCodeResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		URL          string `json:"url"`
		RefreshToken string `json:"refresh_token"`
		Timestamp    int    `json:"timestamp"`
		Code         int    `json:"code"`
		Message      string `json:"message"`
	} `json:"data"`
}

// GenerateQrcode https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/login/login_action/QR.md
func (client *Client) GenerateQrcode() (*GenerateQrCodeResp, error) {
	client.HttpClient = &http.Client{}
	resp, err := client.HttpClient.Get(generateQrCodeUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.ErrUnexpectedStatusCode(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	generateQrCodeResp := &GenerateQrCodeResp{}
	if err = json.Unmarshal(body, generateQrCodeResp); err != nil {
		return nil, err
	}

	if generateQrCodeResp.Code != 0 {
		return nil, errors.StatusError{Code: generateQrCodeResp.Code, Cause: generateQrCodeResp.Message}
	}
	return generateQrCodeResp, nil
}

// PollQrcode https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/login/login_action/QR.md
func (client *Client) PollQrcode(qrcode string) (*PollQrCodeResp, error) {
	client.HttpClient = &http.Client{}

	url := fmt.Sprintf("%s?qrcode_key=%s", pollQrCodeUrl, qrcode)
	resp, err := client.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.ErrUnexpectedStatusCode(resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pollQrCodeResp := &PollQrCodeResp{}
	if err = json.Unmarshal(body, pollQrCodeResp); err != nil {
		return nil, err
	}

	if pollQrCodeResp.Code != 0 {
		return nil, errors.StatusError{Code: pollQrCodeResp.Code, Cause: pollQrCodeResp.Message}
	}
	if pollQrCodeResp.Data.Code == 0 {
		client.readCookieFromHeader(resp.Header)
	}
	return pollQrCodeResp, nil
}

// LoginStatus https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/login/login_action/QR.md#%E6%89%AB%E7%A0%81%E7%99%BB%E5%BD%95web%E7%AB%AF
type LoginStatus int

var (
	LoginSuccess           = LoginStatus(0)
	LoginNotScan           = LoginStatus(86101)
	LoginScanButNotConfirm = LoginStatus(86090)
	LoginExpired           = LoginStatus(86038)
)

// LoginWithQrCode writer is where the qrcode be written
func (client *Client) LoginWithQrCode(writer io.Writer) (<-chan LoginStatus, error) {
	generateQrCodeResp, err := client.GenerateQrcode()
	if err != nil {
		return nil, err
	}

	if err = qrcode.Generate(generateQrCodeResp.Data.URL, qrcode.Low, writer); err != nil {
		return nil, err
	}

	var loginStatus = make(chan LoginStatus)
	go func() {
		defer close(loginStatus)
		var pollQrCodeResp *PollQrCodeResp
		for {
			pollQrCodeResp, err = client.PollQrcode(generateQrCodeResp.Data.QrcodeKey)
			if err != nil {
				loginStatus <- -1
				return
			}
			loginStatus <- LoginStatus(pollQrCodeResp.Data.Code)
			switch pollQrCodeResp.Data.Code {
			case int(LoginSuccess), int(LoginExpired):
				return
			default:
				continue
			}
		}
	}()
	return loginStatus, nil
}
