package main

import (
	"github.com/ljun20160606/di"
	"github.com/urfave/cli"
	"log"
	"os"
)

// 1. 生成leetcode.json
// 2. 根据题目生成题目所需文件
//   (1) 判断文件是否在字典中
//   (2) 抓取描述
//   (3) 根据Problem建立文件
func main() {
	err := di.ConfigLoadFile("./config.yaml", di.YAML)
	if err != nil {
		log.Fatal(err)
	}
	di.Start()
	app := cli.NewApp()
	app.Name = "leetcode"
	app.Version = "0.3.0"
	app.Usage = "leetcode刷题辅助工具"
	app.Commands = append(lc.Commands)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
