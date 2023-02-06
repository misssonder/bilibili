# Bilibili Download
本项目是使用golang编写的bilibili视频下载器

## 安装
```shell
go install github.com/misssonder/bilibili/cmd/bilibilidl
```

## Example
```shell
$ bilibilidl -h 

Bilibili Downloader

Usage:
  bilibilidl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  download    Download bilibili video through url/BVID/AVID.
  help        Help about any command
  info        Show base info of video.
  login       Login bilibili through qrcode (default is $HOME/.bilibili_cookie.txt).

Flags:
  -h, --help      help for bilibilidl
  -v, --verbose   Enable verbose output

Use "bilibilidl [command] --help" for more information about a command.
```
### 视频基本信息
```shell
$ bilibilidl info 'https://www.bilibili.com/video/BV16X4y1g7wT/?spm_id_from=333.934.0.0'

Title:       【才浅】15天花20万元用500克黄金敲数万锤纯手工打造三星堆黄金面具
Author:      才疏学浅的才浅(2200736)
Duration:    717
BvID:        BV16X4y1g7wT
AID:         715024588
Description: 倾家荡产求三连支持！！！请大家帮忙给新系列想个名字，点赞一百万的话制作三星堆黄金权杖，不会真的可以点到一百万吧
bgm:
-Old-B - 【Free Beat】侠之道 、于剑飞 - 01 片头曲 帝陵、AniFace - 夜辞秋江

+-----------------------+------+-----------+----------+-----------+
|         PART          | PAGE |    CID    | DURATION | DIMENSION |
+-----------------------+------+-----------+----------+-----------+
| 三星堆面具字幕修改版3    |    1 | 323723441 |      717 | 1080*1920 |
+-----------------------+------+-----------+----------+-----------+
```
```shell
$  bilibilidl info 'https://www.bilibili.com/bangumi/play/ss33622?from_spmid=666.24.0.0'                                                                            (base) 
https://api.bilibili.com/pgc/view/web/season?season_id=33622

Title:       西游记(已观看1.4亿次)
Duration:    84982000
SeasonID:    33622
Description: 东胜神州的傲来国花果山的一块巨石孕育出了一只灵明石猴（六小龄童 饰），石猴后来拜须菩提为师后习得了七十二变，具有了通天本领，于是占山为王，自称齐天大圣。玉帝派太白金星下凡招安大圣上了天庭，后来大圣因为嫌玉帝赐封的官职太低师徒踏上了取经的路途。

+------------+--------------+-----------+---------+-----------+-----------+
|   TITLE    |     BVID     |    CID    |   AID   | DURATION  | DIMENSION |
+------------+--------------+-----------+---------+-----------+-----------+
| 猴王初问世 | BV1kv411z7F2 | 209571792 | 2668000 | 243343066 | 576*768   |
| 官封弼马温 | BV1NA411q7xW | 209572854 | 2933000 | 328259998 | 576*768   |
| 大圣闹天宫 | BV1vZ4y1p7Jg | 209573621 | 3700000 | 370757035 | 576*768   |
| 困囚五行山 | BV1V5411x7Ty | 209574674 | 2940000 | 455861894 | 576*768   |
| 猴王保唐僧 | BV1Nz411i75F | 209576287 | 3345000 | 200871267 | 576*768   |
| 祸起观音院 | BV1AZ4y1W7MW | 209577244 | 2805000 | 370814224 | 576*768   |
| 计收猪八戒 | BV1AT4y1g7hD | 209577986 | 3310000 | 925818386 | 576*768   |
| 坎途逢三难 | BV1yi4y147N3 | 209579168 | 3263000 | 540763650 | 576*768   |
| 偷吃人参果 | BV1fi4y147XE | 209580032 | 3329000 | 540778911 | 576*768   |
| 三打白骨精 | BV1Hp4y1X7dS | 209580823 | 3185000 | 968357315 | 576*768   |
| 智激美猴王 | BV1sQ4y1P7PP | 209582099 | 3417000 | 710842807 | 576*768   |
| 夺宝莲花洞 | BV1XC4y1h7Zz | 209583092 | 3335000 | 795873129 | 576*768   |
| 除妖乌鸡国 | BV1qZ4y1p7pU | 209584507 | 3788000 | 370757712 | 576*768   |
| 大战红孩儿 | BV1Fp4y1X7ud | 209585634 | 2789000 | 968283683 | 576*768   |
| 斗法降三怪 | BV1az4y1d7mt | 209587679 | 3952000 | 583251656 | 576*768   |
| 趣经女儿国 | BV1sK411s7J5 | 209588655 | 3906000 | 498308615 | 576*768   |
| 三调芭蕉扇 | BV1tT4y1g7k3 | 209589790 | 3329000 | 925867098 | 576*768   |
| 扫塔辨奇冤 | BV1cz4y197ug | 209590866 | 3636000 | 583286723 | 576*768   |
| 误入小雷音 | BV1Bg4y1B7tA | 209592191 | 3680000 | 838268614 | 576*768   |
| 孙猴巧行医 | BV1tg4y1i7oa | 209593494 | 3615000 | 838337538 | 576*768   |
| 错坠盘丝洞 | BV1Lv411z7hS | 209594642 | 3745000 | 243290099 | 576*768   |
| 四探无底洞 | BV1HA411q7tJ | 209596266 | 4055000 | 328280878 | 576*768   |
| 传艺玉华洲 | BV1Mp4y1X7eo | 209597635 | 3813000 | 968312146 | 576*768   |
| 天竺收玉兔 | BV18V411C7xu | 209598791 | 3066000 | 413285628 | 576*768   |
| 波生极乐天 | BV18v411z74W | 209599683 | 3378000 | 243332804 | 576*768   |
+------------+--------------+-----------+---------+-----------+-----------+

```
### 下载视频
- [x] 下载用户上传视频（通过输入BV号或者网址）
- [x] 下载剧集（通过输入剧集网址）

![](images/example_download.gif)
![](images/example_download_season.gif)
## Inspired
- [https://github.com/SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)
- [https://github.com/kkdai/youtube](https://github.com/kkdai/youtube)