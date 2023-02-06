package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	bilibili "github.com/misssonder/bilibili/pkg/client"
	"github.com/misssonder/bilibili/pkg/video"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type Page struct {
	CID       int64
	Page      int
	Part      string
	Duration  int64
	Dimension Dimension
}
type Dimension struct {
	Height int
	Width  int
}

type VideoInfo struct {
	BvID        string
	AID         int
	Title       string
	Author      string
	Duration    int64
	PublishTime string
	CreateTime  string
	Description string
	Pages       []Page
}

type SeasonInfo struct {
	SeasonID    int
	Title       string
	Duration    int64
	Description string
	Episodes    []Episode
}

type Episode struct {
	BvID      string
	AID       int
	CID       int64
	Title     string
	Duration  int
	Dimension Dimension
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show base info of video.",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := checkOutputFormat(); err != nil {
			return err
		}
		return login()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if video.IsEpID(args[0]) || video.IsSSID(args[0]) {
			seasonInfo, err := getSeasonInfo(args[0])
			exitOnError(err)
			exitOnError(writeOutput(os.Stdout, seasonInfo, func(w io.Writer) {
				writeSeasonInfoOutput(w, seasonInfo)
			}))
		} else {
			videoInfo, err := getVideoInfo(args[0])
			exitOnError(err)
			exitOnError(writeOutput(os.Stdout, videoInfo, func(w io.Writer) {
				writeVideoInfoOutput(w, videoInfo)
			}))
		}

	},
}

func getSeasonInfo(id string) (seasonInfo *SeasonInfo, err error) {
	var info *bilibili.SeasonSectionResp
	if video.IsSSID(id) {
		ssID, err := video.ExtractSSID(id)
		if err != nil {
			return nil, err
		}
		info, err = client.SeasonSection(ssID, "")
	} else {
		epID, err := video.ExtractEpID(id)
		if err != nil {
			return nil, err
		}
		info, err = client.SeasonSection("", epID)
	}
	if err != nil {
		return nil, err
	}
	seasonInfo = &SeasonInfo{
		SeasonID:    info.Result.SeasonID,
		Title:       fmt.Sprintf("%s(%s)", info.Result.Title, info.Result.Subtitle),
		Description: info.Result.Evaluate,
		Episodes:    make([]Episode, 0),
	}
	for _, episode := range info.Result.Episodes {
		page := Episode{
			BvID:     episode.Bvid,
			CID:      int64(episode.Cid),
			AID:      episode.Aid,
			Duration: episode.Duration,
			Title:    episode.LongTitle,
		}
		if episode.Dimension.Rotate != 0 {
			page.Dimension.Height = episode.Dimension.Width
			page.Dimension.Width = episode.Dimension.Height
		} else {
			page.Dimension.Height = episode.Dimension.Height
			page.Dimension.Width = episode.Dimension.Width
		}
		seasonInfo.Episodes = append(seasonInfo.Episodes, page)
		seasonInfo.Duration += int64(episode.Duration)
	}
	return
}

func getVideoInfo(id string) (videoInfo *VideoInfo, err error) {
	id, err = video.ExtractBvID(id)
	if err != nil {
		return nil, err
	}
	info, err := client.GetVideoInfo(id)
	if err != nil {
		return nil, err
	}
	videoInfo = &VideoInfo{
		BvID:        info.Data.Bvid,
		AID:         info.Data.Aid,
		Title:       info.Data.Title,
		Author:      fmt.Sprintf("%s(%d)", info.Data.Owner.Name, info.Data.Owner.Mid),
		PublishTime: time.Unix(int64(info.Data.Pubdate), 0).Format(time.RFC3339),
		CreateTime:  time.Unix(int64(info.Data.Ctime), 0).Format(time.RFC3339),
		Description: info.Data.Desc,
		Pages:       make([]Page, 0),
	}
	for _, p := range info.Data.Pages {
		page := Page{
			CID:      int64(p.Cid),
			Duration: int64(p.Duration),
			Part:     p.Part,
			Page:     p.Page,
		}
		if p.Dimension.Rotate != 0 {
			page.Dimension.Height = p.Dimension.Width
			page.Dimension.Width = p.Dimension.Height
		} else {
			page.Dimension.Height = p.Dimension.Height
			page.Dimension.Width = p.Dimension.Width
		}
		videoInfo.Pages = append(videoInfo.Pages, page)
		videoInfo.Duration += int64(p.Duration)
	}
	return
}

func writeVideoInfoOutput(w io.Writer, info *VideoInfo) {
	fmt.Println("Title:      ", info.Title)
	fmt.Println("Author:     ", info.Author)
	fmt.Println("Duration:   ", info.Duration)
	fmt.Println("BvID:       ", info.BvID)
	fmt.Println("AID:        ", info.AID)
	fmt.Println("Description:", info.Description)
	fmt.Println()
	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{
		"part",
		"page",
		"cid",
		"duration",
		"dimension",
	})
	for _, page := range info.Pages {
		table.Append([]string{
			page.Part,
			strconv.Itoa(page.Page),
			strconv.Itoa(int(page.CID)),
			strconv.Itoa(int(page.Duration)),
			fmt.Sprintf("%d*%d", page.Dimension.Height, page.Dimension.Width),
		})
	}
	table.Render()
}

func writeSeasonInfoOutput(w io.Writer, info *SeasonInfo) {
	fmt.Println("Title:      ", info.Title)
	fmt.Println("Duration:   ", info.Duration)
	fmt.Println("SeasonID:   ", info.SeasonID)
	fmt.Println("Description:", info.Description)
	fmt.Println()
	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{
		"title",
		"bvid",
		"cid",
		"aid",
		"duration",
		"dimension",
	})
	for _, episode := range info.Episodes {
		table.Append([]string{
			episode.Title,
			episode.BvID,
			strconv.Itoa(int(episode.CID)),
			strconv.Itoa(episode.Duration),
			strconv.Itoa(episode.AID),
			fmt.Sprintf("%d*%d", episode.Dimension.Height, episode.Dimension.Width),
		})
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(infoCmd)
	addFormatFlag(infoCmd.Flags())
}
