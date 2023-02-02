package main

import (
	"io"
	"net/http"
	"os"

	bilibili "github.com/misssonder/bilibili/pkg/client"
	"github.com/misssonder/bilibili/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

// TODO add quality parameter
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "",
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := checkOutputFormat(); err != nil {
			return err
		}
		return login()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := &bilibili.Client{}
		info, err := client.GetVideoInfo(args[0])
		if err != nil {
			exitOnError(err)
		}

		playUrl, err := client.PlayUrl(args[0], int64(info.Data.Pages[0].Cid), 120, bilibili.HDR|bilibili.Dash)
		if err != nil {
			exitOnError(err)
		}

		file, err := os.OpenFile("video.mp4", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			exitOnError(err)
		}
		defer file.Close()

		if err := download(playUrl.Data.Dash.Video[0].BaseURL, file); err != nil {
			exitOnError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func download(url string, writer io.Writer) error {
	cli := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("referer", "https://www.bilibili.com")
	resp, err := cli.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.ErrUnexpectedStatusCode(resp.StatusCode)
	}
	progress := mpb.New(mpb.WithWidth(64))
	bar := progress.AddBar(
		resp.ContentLength,

		mpb.PrependDecorators(
			decor.CountersKibiByte("% .2f / % .2f"),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.EwmaETA(decor.ET_STYLE_GO, 90),
			decor.Name(" | "),
			decor.EwmaSpeed(decor.UnitKiB, "% .2f", 60),
		),
	)
	reader := bar.ProxyReader(resp.Body)
	_, err = io.Copy(writer, reader)
	progress.Wait()
	return err
}
