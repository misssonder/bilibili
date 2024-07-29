package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/misssonder/bilibili/pkg/errors"
	"github.com/misssonder/bilibili/pkg/video"
)

const (
	videoInfoUrl = "https://api.bilibili.com/x/web-interface/view"
	playUrl      = "https://api.bilibili.com/x/player/playurl"
	playUrlV2    = "https://api.bilibili.com/pgc/player/web/v2/playurl"
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
		Durl              []Durl   `json:"durl"`
		Dash              Dash     `json:"dash"`
		SupportFormats    []struct {
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

type PlayUrlV2Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		PlayCheck struct {
			CanPlay    bool   `json:"can_play"`
			PlayDetail string `json:"play_detail"`
		} `json:"play_check"`
		PlayViewBusinessInfo struct {
			EpisodeInfo struct {
				Aid                   int    `json:"aid"`
				Bvid                  string `json:"bvid"`
				Cid                   int    `json:"cid"`
				DeliveryFragmentVideo bool   `json:"delivery_fragment_video"`
				EpID                  int    `json:"ep_id"`
				EpStatus              int    `json:"ep_status"`
				Interaction           struct {
					Interaction bool `json:"interaction"`
				} `json:"interaction"`
				LongTitle string `json:"long_title"`
				Title     string `json:"title"`
			} `json:"episode_info"`
			SeasonInfo struct {
				SeasonID   int `json:"season_id"`
				SeasonType int `json:"season_type"`
			} `json:"season_info"`
			UserStatus struct {
				FollowInfo struct {
					Follow       int `json:"follow"`
					FollowStatus int `json:"follow_status"`
				} `json:"follow_info"`
				IsLogin int `json:"is_login"`
				PayInfo struct {
					PayCheck    int `json:"pay_check"`
					PayPackPaid int `json:"pay_pack_paid"`
					Sponsor     int `json:"sponsor"`
				} `json:"pay_info"`
				VipInfo struct {
					DueDate int64 `json:"due_date"`
					RealVip bool  `json:"real_vip"`
					Status  int   `json:"status"`
					Type    int   `json:"type"`
				} `json:"vip_info"`
				WatchProgress struct {
					LastEpID    int    `json:"last_ep_id"`
					LastEpIndex string `json:"last_ep_index"`
					LastTime    int    `json:"last_time"`
				} `json:"watch_progress"`
			} `json:"user_status"`
		} `json:"play_view_business_info"`
		Rights struct {
			PayCheck int `json:"pay_check"`
		} `json:"rights"`
		VideoInfo struct {
			Code         int    `json:"code"`
			SeekParam    string `json:"seek_param"`
			IsPreview    int    `json:"is_preview"`
			Fnval        int    `json:"fnval"`
			VideoProject bool   `json:"video_project"`
			Fnver        int    `json:"fnver"`
			Type         string `json:"type"`
			Bp           int    `json:"bp"`
			Result       string `json:"result"`
			SeekType     string `json:"seek_type"`
			DrmTechType  int    `json:"drm_tech_type"`
			VipType      int    `json:"vip_type"`
			From         string `json:"from"`
			VideoCodecid int    `json:"video_codecid"`
			RecordInfo   struct {
				RecordIcon string `json:"record_icon"`
				Record     string `json:"record"`
			} `json:"record_info"`
			IsDrm          bool   `json:"is_drm"`
			NoRexcode      int    `json:"no_rexcode"`
			Format         string `json:"format"`
			SupportFormats []struct {
				DisplayDesc    string   `json:"display_desc"`
				SubDescription string   `json:"sub_description"`
				Superscript    string   `json:"superscript"`
				NeedLogin      bool     `json:"need_login,omitempty"`
				Codecs         []string `json:"codecs"`
				Format         string   `json:"format"`
				Description    string   `json:"description"`
				NeedVip        bool     `json:"need_vip,omitempty"`
				Quality        int      `json:"quality"`
				NewDescription string   `json:"new_description"`
			} `json:"support_formats"`
			Message      string `json:"message"`
			Quality      int    `json:"quality"`
			Timelength   int    `json:"timelength"`
			HasPaid      bool   `json:"has_paid"`
			DrmType      string `json:"drm_type"`
			VipStatus    int    `json:"vip_status"`
			Durl         []Durl `json:"durl"`
			Dash         Dash   `json:"dash"`
			ClipInfoList []struct {
				MaterialNo int    `json:"materialNo"`
				Start      int    `json:"start"`
				End        int    `json:"end"`
				ToastText  string `json:"toastText"`
				ClipType   string `json:"clipType"`
			} `json:"clip_info_list"`
			Status int `json:"status"`
		} `json:"video_info"`
		ViewInfo struct {
			AiRepairQnTrialInfo struct {
				TrialAble bool `json:"trial_able"`
			} `json:"ai_repair_qn_trial_info"`
			EndPage struct {
				Hide bool `json:"hide"`
			} `json:"end_page"`
			ExtToast struct {
				VipDefinitionRemind struct {
					Icon          string `json:"icon"`
					ShowStyleType int    `json:"showStyleType"`
					ToastText     struct {
						Text string `json:"text"`
					} `json:"toast_text"`
				} `json:"VIP_DEFINITION_REMIND"`
				VipDefinitionGuide struct {
					Button struct {
						Text string `json:"text"`
					} `json:"button"`
					Icon          string `json:"icon"`
					ShowStyleType int    `json:"showStyleType"`
					ToastText     struct {
						Text string `json:"text"`
					} `json:"toast_text"`
				} `json:"VIP_DEFINITION_GUIDE"`
			} `json:"ext_toast"`
			PayTip struct {
				AngleStyle        int    `json:"angle_style"`
				BgDayColor        string `json:"bg_day_color"`
				BgLineColor       string `json:"bg_line_color"`
				BgNightColor      string `json:"bg_night_color"`
				BgNightLineColor  string `json:"bg_night_line_color"`
				GiantScreenImg    string `json:"giant_screen_img"`
				Icon              string `json:"icon"`
				Img               string `json:"img"`
				JumpType          string `json:"jump_type"`
				Link              string `json:"link"`
				OrderReportParams struct {
					TipsRepeatKey string `json:"tips_repeat_key"`
					EpStatus      string `json:"ep_status"`
					ExpTag        string `json:"exp_tag"`
					SeasonID      string `json:"season_id"`
					SeasonStatus  string `json:"season_status"`
					EpID          string `json:"ep_id"`
					MaterialType  string `json:"material_type"`
					SeasonType    string `json:"season_type"`
					VipType       string `json:"vip_type"`
					VipStatus     string `json:"vip_status"`
					UnitID        string `json:"unit_id"`
					TipsID        string `json:"tips_id"`
					RequestID     string `json:"request_id"`
					ExpGroupTag   string `json:"exp_group_tag"`
					PositionID    string `json:"position_id"`
				} `json:"order_report_params"`
				PcLink string `json:"pc_link"`
				Report struct {
					ClickEventID string `json:"clickEventId"`
					Extend       string `json:"extend"`
					ShowEventID  string `json:"showEventId"`
				} `json:"report"`
				ReportType     int    `json:"report_type"`
				ShowType       int    `json:"show_type"`
				TextNightColor string `json:"textNightColor"`
				TextColor      string `json:"text_color"`
				Title          string `json:"title"`
				Type           int    `json:"type"`
				URLOpenType    int    `json:"url_open_type"`
				ViewStartTime  int    `json:"view_start_time"`
			} `json:"pay_tip"`
			QnTrialInfo struct {
				TrialAble bool `json:"trial_able"`
			} `json:"qn_trial_info"`
			Report struct {
				EpID         string `json:"ep_id"`
				EpStatus     string `json:"ep_status"`
				SeasonID     string `json:"season_id"`
				SeasonStatus string `json:"season_status"`
				SeasonType   string `json:"season_type"`
				VipStatus    string `json:"vip_status"`
				VipType      string `json:"vip_type"`
			} `json:"report"`
		} `json:"view_info"`
	} `json:"result"`
}

type Dash struct {
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
}

type Durl struct {
	Size      int      `json:"size"`
	Ahead     string   `json:"ahead"`
	Length    int      `json:"length"`
	Vhead     string   `json:"vhead"`
	BackupURL []string `json:"backup_url"`
	URL       string   `json:"url"`
	Order     int      `json:"order"`
	Md5       string   `json:"md5"`
}

func (client *Client) GetVideoInfo(id string) (*VideoInfoResp, error) {
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
	FnvalMP4  Fnval = 1
	FnvalDash Fnval = 16
	FnvalHDR  Fnval = 64
	Fnval4K   Fnval = 128

	FnvalAudio64K  Fnval = 30216
	FnvalAudio132K Fnval = 30232
	FnvalAudio192K Fnval = 30280
)

type Qn int64

func (qn Qn) String() string {
	switch qn {
	case Qn240P:
		return "240P"
	case Qn360P:
		return "360P"
	case Qn480P:
		return "480P"
	case Qn720P:
		return "720P"
	case Qn720P60:
		return "720P60"
	case Qn1080P:
		return "1080P"
	case Qn1080PPlus:
		return "1080P+"
	case Qn1080P60:
		return "1080P60"
	case Qn4k:
		return "4K"
	case QnAudio64K:
		return "64K"
	case QnAudio132K:
		return "132K"
	case QnAudio192K:
		return "192K"
	case QnAudioDolby:
		return "Dolby"
	case QnAudioHiRes:
		return "Hi-Res"
	default:
		return ""
	}
}

const (
	Qn240P      Qn = 6
	Qn360P      Qn = 16
	Qn480P      Qn = 32
	Qn720P      Qn = 64
	Qn720P60    Qn = 74
	Qn1080P     Qn = 80
	Qn1080PPlus Qn = 112
	Qn1080P60   Qn = 116
	Qn4k        Qn = 120

	QnAudio64K   Qn = 30216
	QnAudio132K  Qn = 30232
	QnAudio192K  Qn = 30280
	QnAudioDolby Qn = 30250
	QnAudioHiRes Qn = 30251
)

func (client *Client) PlayUrl(bvid string, cid int64, qn Qn, fnval Fnval) (*PlayUrlResp, error) {
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

func (client *Client) PlayUrlV2(epid int64, qn Qn, fnval Fnval) (*PlayUrlV2Resp, error) {
	url := fmt.Sprintf("%s?ep_id=%d&qn=%d&fnval=%d&fnver=0&fourk=1&support_multi_audio=true&gaia_source=&is_main_page=true&need_fragment=true&isGaiaAvoided=false&voice_balance=1&drm_tech_type=2", playUrlV2, epid, qn, fnval)
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("referer", "https://www.bilibili.com/")
	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.ErrUnexpectedStatusCode(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	playUrlResp := &PlayUrlV2Resp{}
	if err = json.Unmarshal(body, playUrlResp); err != nil {
		return nil, err
	}
	if playUrlResp.Code != 0 {
		return nil, errors.StatusError{Code: playUrlResp.Code, Cause: playUrlResp.Message}
	}
	return playUrlResp, nil
}
