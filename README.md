# Bilibili Download
本项目是使用golang编写的bilibili视频下载器

# 安装
```shell
go install github.com/misssonder/bilibili/cmd/bilibilidl
```

# Example
## 视频基本信息
```shell
bilibilidl info 'https://www.bilibili.com/video/BV16X4y1g7wT/?spm_id_from=333.934.0.0'

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
## 下载视频
