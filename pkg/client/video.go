package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/misssonder/bilibili/pkg/errors"
	"github.com/misssonder/bilibili/pkg/video"
)

const (
	videoInfoUrl = "https://api.bilibili.com/x/web-interface/view"
	playUrl      = "https://api.bilibili.com/x/player/playurl"
)

type VideoInfoResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizID   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionID int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			ArgueMsg   string `json:"argue_msg"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Premiere           interface{} `json:"premiere"`
		TeenageMode        int         `json:"teenage_mode"`
		IsChargeableSeason bool        `json:"is_chargeable_season"`
		IsStory            bool        `json:"is_story"`
		NoCache            bool        `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		Staff []struct {
			Mid   int    `json:"mid"`
			Title string `json:"title"`
			Name  string `json:"name"`
			Face  string `json:"face"`
			Vip   struct {
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
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official"`
			Follower   int `json:"follower"`
			LabelStyle int `json:"label_style"`
		} `json:"staff"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			URLImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
			Honor []struct {
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				Desc               string `json:"desc"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
		LikeIcon string `json:"like_icon"`
	} `json:"data"`
}

type PlayUrlResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Durl              []struct {
			Order     int      `json:"order"`
			Length    int      `json:"length"`
			Size      int      `json:"size"`
			Ahead     string   `json:"ahead"`
			Vhead     string   `json:"vhead"`
			URL       string   `json:"url"`
			BackupURL []string `json:"backup_url"`
		} `json:"durl"`
		Dash struct {
			Duration      int     `json:"duration"`
			MinBufferTime float64 `json:"min_buffer_time"`
			Video         []struct {
				ID           int      `json:"id"`
				BaseURL      string   `json:"base_url"`
				BackupURL    []string `json:"backup_url"`
				Bandwidth    int      `json:"bandwidth"`
				MimeType     string   `json:"mime_type"`
				Codecs       string   `json:"codecs"`
				Width        int      `json:"width"`
				Height       int      `json:"height"`
				FrameRate    string   `json:"frame_rate"`
				Sar          string   `json:"sar"`
				StartWithSap int      `json:"start_with_sap"`
				SegmentBase  struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				ID           int      `json:"id"`
				BaseURL      string   `json:"base_url"`
				BackupURL    []string `json:"backup_url"`
				Bandwidth    int      `json:"bandwidth"`
				MimeType     string   `json:"mime_type"`
				Codecs       string   `json:"codecs"`
				Width        int      `json:"width"`
				Height       int      `json:"height"`
				FrameRate    string   `json:"frame_rate"`
				Sar          string   `json:"sar"`
				StartWithSap int      `json:"start_with_sap"`
				SegmentBase  struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"audio"`
			Dolby struct {
				Type  int         `json:"type"`
				Audio interface{} `json:"audio"`
			} `json:"dolby"`
			Flac interface{} `json:"flac"`
		} `json:"dash"`
		SupportFormats []struct {
			Quality        int         `json:"quality"`
			Format         string      `json:"format"`
			NewDescription string      `json:"new_description"`
			DisplayDesc    string      `json:"display_desc"`
			Superscript    string      `json:"superscript"`
			Codecs         interface{} `json:"codecs"`
		} `json:"support_formats"`
		HighFormat   interface{} `json:"high_format"`
		LastPlayTime int         `json:"last_play_time"`
		LastPlayCid  int         `json:"last_play_cid"`
	} `json:"data"`
}

func (client *Client) GetVideoInfo(id string) (*VideoInfoResp, error) {
	id, err := video.ExtractBvID(id)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?bvid=%s", videoInfoUrl, id)
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, url, nil)
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

	videoInfoResp := &VideoInfoResp{}
	if err = json.Unmarshal(body, videoInfoResp); err != nil {
		return nil, err
	}
	if videoInfoResp.Code != 0 {
		return nil, errors.StatusError{Code: videoInfoResp.Code, Cause: videoInfoResp.Message}
	}

	return videoInfoResp, nil
}

// Fnval https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/video/videostream_url.md#fnval%E8%A7%86%E9%A2%91%E6%B5%81%E6%A0%BC%E5%BC%8F%E6%A0%87%E8%AF%86
type Fnval int64

const (
	MP4  Fnval = 1
	Dash Fnval = 16
	HDR  Fnval = 64
	_4K  Fnval = 128
)

func (client *Client) PlayUrl(bvid string, cid int64, qn int, fnval Fnval) (*PlayUrlResp, error) {
	id, err := video.ExtractBvID(bvid)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?bvid=%s&cid=%d&qn=%d&fourk=1&fnval=%d", playUrl, id, cid, qn, fnval)
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, url, nil)
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

	playUrlResp := &PlayUrlResp{}
	if err = json.Unmarshal(body, playUrlResp); err != nil {
		return nil, err
	}
	if playUrlResp.Code != 0 {
		return nil, errors.StatusError{Code: playUrlResp.Code, Cause: playUrlResp.Message}
	}

	return playUrlResp, nil
}
