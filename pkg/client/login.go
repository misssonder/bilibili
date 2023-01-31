package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/misssonder/bilibili/pkg/errors"
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
	return pollQrCodeResp, nil
}
