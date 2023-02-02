package main

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	file, err := os.OpenFile("video.mp4", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	err = downloadMedia("https://cn-hbwh-cm-01-09.bilivideo.com/upgcxcode/00/39/983613900/983613900_sr1-1-100035.m4s?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1675330209&gen=playurlv2&os=bcache&oi=2029904162&trid=0000d71a9e39d2154bd9adf84d593118ac1bu&mid=35988173&platform=pc&upsig=749f0235988fd8ba079349a3e920457f&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,mid,platform&cdnid=10205&bvc=vod&nettype=0&orderid=0,3&buvid=&build=0&agrr=1&bw=1419861&logo=80000000", file)
	if err != nil {
		t.Error(err)
		return
	}
}
