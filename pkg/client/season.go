package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/misssonder/bilibili/pkg/errors"
)

const (
	seasonSectionInfoUrl = "https://api.bilibili.com/pgc/view/web/season"
)

type SeasonSectionResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Activity struct {
			HeadBgURL string `json:"head_bg_url"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
		} `json:"activity"`
		Alias string `json:"alias"`
		Areas []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"areas"`
		BkgCover string `json:"bkg_cover"`
		Cover    string `json:"cover"`
		Episodes []struct {
			Aid       int    `json:"aid"`
			Badge     string `json:"badge"`
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			BadgeType int    `json:"badge_type"`
			Bvid      string `json:"bvid"`
			Cid       int    `json:"cid"`
			Cover     string `json:"cover"`
			Dimension struct {
				Height int `json:"height"`
				Rotate int `json:"rotate"`
				Width  int `json:"width"`
			} `json:"dimension"`
			Duration    int    `json:"duration"`
			From        string `json:"from"`
			ID          int    `json:"id"`
			IsViewHide  bool   `json:"is_view_hide"`
			Link        string `json:"link"`
			LongTitle   string `json:"long_title"`
			PubTime     int    `json:"pub_time"`
			Pv          int    `json:"pv"`
			ReleaseDate string `json:"release_date"`
			Rights      struct {
				AllowDemand   int `json:"allow_demand"`
				AllowDm       int `json:"allow_dm"`
				AllowDownload int `json:"allow_download"`
				AreaLimit     int `json:"area_limit"`
			} `json:"rights"`
			ShareCopy string `json:"share_copy"`
			ShareURL  string `json:"share_url"`
			ShortLink string `json:"short_link"`
			Skip      struct {
				Ed struct {
					End   int `json:"end"`
					Start int `json:"start"`
				} `json:"ed"`
				Op struct {
					End   int `json:"end"`
					Start int `json:"start"`
				} `json:"op"`
			} `json:"skip"`
			Status   int    `json:"status"`
			Subtitle string `json:"subtitle"`
			Title    string `json:"title"`
			Vid      string `json:"vid"`
		} `json:"episodes"`
		Evaluate string `json:"evaluate"`
		Freya    struct {
			BubbleDesc    string `json:"bubble_desc"`
			BubbleShowCnt int    `json:"bubble_show_cnt"`
			IconShow      int    `json:"icon_show"`
		} `json:"freya"`
		JpTitle string `json:"jp_title"`
		Link    string `json:"link"`
		MediaID int    `json:"media_id"`
		Mode    int    `json:"mode"`
		NewEp   struct {
			Desc  string `json:"desc"`
			ID    int    `json:"id"`
			IsNew int    `json:"is_new"`
			Title string `json:"title"`
		} `json:"new_ep"`
		Positive struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"positive"`
		Publish struct {
			IsFinish      int    `json:"is_finish"`
			IsStarted     int    `json:"is_started"`
			PubTime       string `json:"pub_time"`
			PubTimeShow   string `json:"pub_time_show"`
			UnknowPubDate int    `json:"unknow_pub_date"`
			Weekday       int    `json:"weekday"`
		} `json:"publish"`
		Rating struct {
			Count int     `json:"count"`
			Score float64 `json:"score"`
		} `json:"rating"`
		Record string `json:"record"`
		Rights struct {
			AllowBp         int    `json:"allow_bp"`
			AllowBpRank     int    `json:"allow_bp_rank"`
			AllowDownload   int    `json:"allow_download"`
			AllowReview     int    `json:"allow_review"`
			AreaLimit       int    `json:"area_limit"`
			BanAreaShow     int    `json:"ban_area_show"`
			CanWatch        int    `json:"can_watch"`
			Copyright       string `json:"copyright"`
			ForbidPre       int    `json:"forbid_pre"`
			FreyaWhite      int    `json:"freya_white"`
			IsCoverShow     int    `json:"is_cover_show"`
			IsPreview       int    `json:"is_preview"`
			OnlyVipDownload int    `json:"only_vip_download"`
			Resource        string `json:"resource"`
			WatchPlatform   int    `json:"watch_platform"`
		} `json:"rights"`
		SeasonID    int    `json:"season_id"`
		SeasonTitle string `json:"season_title"`
		Seasons     []struct {
			Badge     string `json:"badge"`
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			BadgeType           int    `json:"badge_type"`
			Cover               string `json:"cover"`
			HorizontalCover1610 string `json:"horizontal_cover_1610"`
			HorizontalCover169  string `json:"horizontal_cover_169"`
			MediaID             int    `json:"media_id"`
			NewEp               struct {
				Cover     string `json:"cover"`
				ID        int    `json:"id"`
				IndexShow string `json:"index_show"`
			} `json:"new_ep"`
			SeasonID    int    `json:"season_id"`
			SeasonTitle string `json:"season_title"`
			SeasonType  int    `json:"season_type"`
			Stat        struct {
				Favorites    int `json:"favorites"`
				SeriesFollow int `json:"series_follow"`
				Views        int `json:"views"`
			} `json:"stat"`
		} `json:"seasons"`
		Section []struct {
			Attr       int           `json:"attr"`
			EpisodeID  int           `json:"episode_id"`
			EpisodeIds []interface{} `json:"episode_ids"`
			Episodes   []struct {
				Aid       int    `json:"aid"`
				Badge     string `json:"badge"`
				BadgeInfo struct {
					BgColor      string `json:"bg_color"`
					BgColorNight string `json:"bg_color_night"`
					Text         string `json:"text"`
				} `json:"badge_info"`
				BadgeType int    `json:"badge_type"`
				Bvid      string `json:"bvid"`
				Cid       int    `json:"cid"`
				Cover     string `json:"cover"`
				Dimension struct {
					Height int `json:"height"`
					Rotate int `json:"rotate"`
					Width  int `json:"width"`
				} `json:"dimension"`
				Duration    int    `json:"duration"`
				From        string `json:"from"`
				ID          int    `json:"id"`
				IsViewHide  bool   `json:"is_view_hide"`
				Link        string `json:"link"`
				LongTitle   string `json:"long_title"`
				PubTime     int    `json:"pub_time"`
				Pv          int    `json:"pv"`
				ReleaseDate string `json:"release_date"`
				Rights      struct {
					AllowDemand   int `json:"allow_demand"`
					AllowDm       int `json:"allow_dm"`
					AllowDownload int `json:"allow_download"`
					AreaLimit     int `json:"area_limit"`
				} `json:"rights"`
				ShareCopy string `json:"share_copy"`
				ShareURL  string `json:"share_url"`
				ShortLink string `json:"short_link"`
				Skip      struct {
					Ed struct {
						End   int `json:"end"`
						Start int `json:"start"`
					} `json:"ed"`
					Op struct {
						End   int `json:"end"`
						Start int `json:"start"`
					} `json:"op"`
				} `json:"skip"`
				Stat struct {
					Coin     int `json:"coin"`
					Danmakus int `json:"danmakus"`
					Likes    int `json:"likes"`
					Play     int `json:"play"`
					Reply    int `json:"reply"`
				} `json:"stat"`
				Status   int    `json:"status"`
				Subtitle string `json:"subtitle"`
				Title    string `json:"title"`
				Vid      string `json:"vid"`
			} `json:"episodes"`
			ID    int    `json:"id"`
			Title string `json:"title"`
			Type  int    `json:"type"`
		} `json:"section"`
		Series struct {
			DisplayType int    `json:"display_type"`
			SeriesID    int    `json:"series_id"`
			SeriesTitle string `json:"series_title"`
		} `json:"series"`
		ShareCopy     string `json:"share_copy"`
		ShareSubTitle string `json:"share_sub_title"`
		ShareURL      string `json:"share_url"`
		Show          struct {
			WideScreen int `json:"wide_screen"`
		} `json:"show"`
		ShowSeasonType int    `json:"show_season_type"`
		SquareCover    string `json:"square_cover"`
		Stat           struct {
			Coins     int `json:"coins"`
			Danmakus  int `json:"danmakus"`
			Favorite  int `json:"favorite"`
			Favorites int `json:"favorites"`
			Likes     int `json:"likes"`
			Reply     int `json:"reply"`
			Share     int `json:"share"`
			Views     int `json:"views"`
		} `json:"stat"`
		Status   int    `json:"status"`
		Subtitle string `json:"subtitle"`
		Title    string `json:"title"`
		Total    int    `json:"total"`
		Type     int    `json:"type"`
		UpInfo   struct {
			Avatar             string `json:"avatar"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			Follower           int    `json:"follower"`
			IsFollow           int    `json:"is_follow"`
			Mid                int    `json:"mid"`
			NicknameColor      string `json:"nickname_color"`
			Pendant            struct {
				Image string `json:"image"`
				Name  string `json:"name"`
				Pid   int    `json:"pid"`
			} `json:"pendant"`
			ThemeType  int    `json:"theme_type"`
			Uname      string `json:"uname"`
			VerifyType int    `json:"verify_type"`
			VipLabel   struct {
				BgColor     string `json:"bg_color"`
				BgStyle     int    `json:"bg_style"`
				BorderColor string `json:"border_color"`
				Text        string `json:"text"`
				TextColor   string `json:"text_color"`
			} `json:"vip_label"`
			VipStatus int `json:"vip_status"`
			VipType   int `json:"vip_type"`
		} `json:"up_info"`
		UserStatus struct {
			AreaLimit    int `json:"area_limit"`
			BanAreaShow  int `json:"ban_area_show"`
			Follow       int `json:"follow"`
			FollowStatus int `json:"follow_status"`
			Login        int `json:"login"`
			Pay          int `json:"pay"`
			PayPackPaid  int `json:"pay_pack_paid"`
			Sponsor      int `json:"sponsor"`
		} `json:"user_status"`
	} `json:"result"`
}

func (client *Client) SeasonSection(ssID string, epID string) (*SeasonSectionResp, error) {
	u, err := url.Parse(seasonSectionInfoUrl)
	if err != nil {
		return nil, err
	}
	values := u.Query()
	if len(ssID) != 0 {
		values.Set("season_id", ssID)
	}
	if len(epID) != 0 {
		values.Set("ep_id", epID)
	}
	u.RawQuery = values.Encode()
	println(u.String())
	client.HttpClient = &http.Client{}
	request, err := client.newCookieRequest(http.MethodGet, u.String(), nil)
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

	seasonSectionResp := &SeasonSectionResp{}
	if err = json.Unmarshal(body, seasonSectionResp); err != nil {
		return nil, err
	}
	if seasonSectionResp.Code != 0 {
		return nil, errors.StatusError{Code: seasonSectionResp.Code, Cause: seasonSectionResp.Message}
	}

	return seasonSectionResp, nil
}
