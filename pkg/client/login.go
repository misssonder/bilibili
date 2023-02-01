package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/misssonder/bilibili/pkg/errors"
	"github.com/misssonder/bilibili/pkg/qrcode"
)

const (
	generateQrCodeUrl = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	pollQrCodeUrl     = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
	navInfoUrl        = "https://api.bilibili.com/x/web-interface/nav"
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

type NavInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		IsLogin       bool   `json:"isLogin"`
		EmailVerified int    `json:"email_verified"`
		Face          string `json:"face"`
		FaceNft       int    `json:"face_nft"`
		FaceNftType   int    `json:"face_nft_type"`
		LevelInfo     struct {
			CurrentLevel int    `json:"current_level"`
			CurrentMin   int    `json:"current_min"`
			CurrentExp   int    `json:"current_exp"`
			NextExp      string `json:"next_exp"`
		} `json:"level_info"`
		Mid            int     `json:"mid"`
		MobileVerified int     `json:"mobile_verified"`
		Money          float64 `json:"money"`
		Moral          int     `json:"moral"`
		Official       struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		OfficialVerify struct {
			Type int    `json:"type"`
			Desc string `json:"desc"`
		} `json:"officialVerify"`
		Pendant struct {
			Pid               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
		} `json:"pendant"`
		Scores       int    `json:"scores"`
		Uname        string `json:"uname"`
		VipDueDate   int64  `json:"vipDueDate"`
		VipStatus    int    `json:"vipStatus"`
		VipType      int    `json:"vipType"`
		VipPayType   int    `json:"vip_pay_type"`
		VipThemeType int    `json:"vip_theme_type"`
		VipLabel     struct {
			Path                  string `json:"path"`
			Text                  string `json:"text"`
			LabelTheme            string `json:"label_theme"`
			TextColor             string `json:"text_color"`
			BgStyle               int    `json:"bg_style"`
			BgColor               string `json:"bg_color"`
			BorderColor           string `json:"border_color"`
			UseImgLabel           bool   `json:"use_img_label"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
		} `json:"vip_label"`
		VipAvatarSubscript int    `json:"vip_avatar_subscript"`
		VipNicknameColor   string `json:"vip_nickname_color"`
		Vip                struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgStyle               int    `json:"bg_style"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				UseImgLabel           bool   `json:"use_img_label"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			TvVipStatus        int    `json:"tv_vip_status"`
			TvVipPayType       int    `json:"tv_vip_pay_type"`
		} `json:"vip"`
		Wallet struct {
			Mid           int `json:"mid"`
			BcoinBalance  int `json:"bcoin_balance"`
			CouponBalance int `json:"coupon_balance"`
			CouponDueTime int `json:"coupon_due_time"`
		} `json:"wallet"`
		HasShop        bool   `json:"has_shop"`
		ShopURL        string `json:"shop_url"`
		AllowanceCount int    `json:"allowance_count"`
		AnswerStatus   int    `json:"answer_status"`
		IsSeniorMember int    `json:"is_senior_member"`
		WbiImg         struct {
			ImgURL string `json:"img_url"`
			SubURL string `json:"sub_url"`
		} `json:"wbi_img"`
		IsJury bool `json:"is_jury"`
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

func (client *Client) NavInfo() (*NavInfoResp, error) {
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, navInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.HttpClient.Do(request)
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

	navInfoResp := &NavInfoResp{}
	if err = json.Unmarshal(body, navInfoResp); err != nil {
		return nil, err
	}

	if navInfoResp.Code != 0 {
		return nil, errors.StatusError{Code: navInfoResp.Code, Cause: navInfoResp.Message}
	}

	return navInfoResp, nil
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
