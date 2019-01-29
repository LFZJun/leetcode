package main

import (
	"errors"
	"github.com/ljun20160606/di"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"strconv"
	"strings"
)

func init() {
	di.Put(new(commandCenter))
}

type commandCenter struct {
	LeetcodeCli *leetcodeCli `di:"*"`
	App         *cli.App     `di:"*"`
}

func (cc *commandCenter) Init() {
	leetcodeCli := cc.LeetcodeCli

	// update
	updateCmd := cli.Command{
		Name:  "update",
		Usage: "更新数据",
	}

	updateCmd.Subcommands = append(updateCmd.Subcommands, cli.Command{
		Name:  "leetcode",
		Usage: "更新leetcode题目表",
		Action: func(c *cli.Context) error {
			var path string
			if c.NArg() > 0 {
				path = c.Args()[0]
			} else {
				path = leetcodeCli.CmdConfig.Lock.Leetcode
			}
			logger.Println("拉取所有leetcode题目至", path)
			lc.WriteCatalog(path)
			return nil
		},
	})

	updateCmd.Subcommands = append(updateCmd.Subcommands, cli.Command{
		Name:  "readme",
		Usage: "重新生成readme",
		Action: func(c *cli.Context) error {
			leetcodeCli.WriteReadme()
			return nil
		},
	})

	cc.App.Commands = append(cc.App.Commands, updateCmd)

	// pull
	cc.App.Commands = append(cc.App.Commands, cli.Command{
		Name:  "pull",
		Usage: "查询本地leetcode.json，然后抓取制定题目",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  flagForce.String(),
				Usage: "先update再从表中查询",
			},
			cli.StringFlag{
				Name:  flagPath.String(),
				Usage: "本地leetcode.json路径",
				Value: leetcodeCli.CmdConfig.Lock.Leetcode,
			},
			cli.BoolFlag{
				Name:  flagInteract.String(),
				Usage: "开启交互模式输入questionID",
			},
		},
		Action: func(c *cli.Context) error {
			var idsStrs []string
			if c.Bool(flagInteract.Flag()) {
				var idsStr = ""
				prompt := &survey.Input{Message: "请输入 QuestionID"}
				_ = survey.AskOne(prompt, &idsStr, nil)
				idsStrs = strings.Split(strings.Trim(idsStr, " "), " ")
				if len(idsStrs) == 0 || len(idsStrs[0]) == 0 {
					return errors.New("请输入 e.g. 1 2 3 4 5")
				}
			} else {
				if c.NArg() < 1 {
					return errors.New("pull [id...]")
				}
				idsStrs = c.Args()
			}

			var ids []int
			for e := range idsStrs {
				id, err := strconv.Atoi(idsStrs[e])
				if err != nil {
					return err
				}
				ids = append(ids, id)
			}

			leetcodeCli.ReadLock()
			leetcodeCli.ReadCatalog(c.String(flagPath.Flag()), c.Bool(flagForce.Flag()))
			leetcodeCli.Gen(ids)
			leetcodeCli.SortItem()
			leetcodeCli.WriteLock()
			leetcodeCli.WriteReadme()
			return nil
		},
	})
}
