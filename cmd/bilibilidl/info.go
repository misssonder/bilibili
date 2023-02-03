package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

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
		info, err := client.GetVideoInfo(args[0])
		exitOnError(err)
		videoInfo := &VideoInfo{
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

		exitOnError(writeOutput(os.Stdout, videoInfo, func(w io.Writer) {
			writeInfoOutput(w, videoInfo)
		}))
	},
}

func writeInfoOutput(w io.Writer, info *VideoInfo) {
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
