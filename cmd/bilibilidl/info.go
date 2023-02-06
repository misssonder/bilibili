package main

import (
	"fmt"
	bilibili "github.com/misssonder/bilibili/pkg/client"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/misssonder/bilibili/pkg/video"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type Page struct {
	Cid       int64
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
	SeasonID    int
	AID         int
	Title       string
	Author      string
	Duration    int64
	PublishTime string
	CreateTime  string
	Description string
	Pages       []Page
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
		var videoInfo = new(VideoInfo)
		switch video.IsSSID(args[0]) || video.IsEpID(args[0]) {
		case true:
			var info *bilibili.SeasonSectionResp
			var err error
			if video.IsSSID(args[0]) {
				ssID, err := video.ExtractSSID(args[0])
				exitOnError(err)
				info, err = client.SeasonSection(ssID, "")
			} else {
				epID, err := video.ExtractEpID(args[0])
				exitOnError(err)
				info, err = client.SeasonSection("", epID)
			}
			exitOnError(err)
			videoInfo = &VideoInfo{
				SeasonID:    info.Result.SeasonID,
				Title:       fmt.Sprintf("%s(%s)", info.Result.Title, info.Result.Subtitle),
				Description: info.Result.Evaluate,
				Pages:       make([]Page, 0),
			}
			for i, e := range info.Result.Episodes {
				page := Page{
					Cid:      int64(e.Cid),
					Duration: int64(e.Duration),
					Part:     e.LongTitle,
					Page:     i + 1,
				}
				if e.Dimension.Rotate != 0 {
					page.Dimension.Height = e.Dimension.Width
					page.Dimension.Width = e.Dimension.Height
				} else {
					page.Dimension.Height = e.Dimension.Height
					page.Dimension.Width = e.Dimension.Width
				}
				videoInfo.Pages = append(videoInfo.Pages, page)
				videoInfo.Duration += int64(e.Duration)
			}
		default:
			info, err := client.GetVideoInfo(args[0])
			exitOnError(err)
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
					Cid:      int64(p.Cid),
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
		}

		exitOnError(writeOutput(os.Stdout, videoInfo, func(w io.Writer) {
			writeInfoOutput(w, videoInfo)
		}))
	},
}

func writeInfoOutput(w io.Writer, info *VideoInfo) {
	fmt.Println("Title:      ", info.Title)
	if info.Author != "" {
		fmt.Println("Author:     ", info.Author)
	}
	fmt.Println("Duration:   ", info.Duration)
	if info.BvID != "" {
		fmt.Println("BvID:       ", info.BvID)
	}
	if info.AID != 0 {
		fmt.Println("AID:        ", info.AID)
	}
	if info.SeasonID != 0 {
		fmt.Println("SeasonID:   ", info.SeasonID)
	}
	fmt.Println("Description:", info.Description)
	fmt.Println()
	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{
		"part",
		"page",
		"cid",
		"duration",
		"Dimension",
	})
	for _, page := range info.Pages {
		table.Append([]string{
			page.Part,
			strconv.Itoa(page.Page),
			strconv.Itoa(int(page.Cid)),
			strconv.Itoa(int(page.Duration)),
			fmt.Sprintf("%d*%d", page.Dimension.Height, page.Dimension.Width),
		})
	}
	table.Render()
}

func init() {
	rootCmd.AddCommand(infoCmd)
	addFormatFlag(infoCmd.Flags())
}
