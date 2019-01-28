package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/ljun20160606/di"
	"github.com/ljun20160606/gox/fs"
	"github.com/ljun20160606/simplehttp"
	"github.com/urfave/cli"
	"gopkg.in/go-playground/pool.v3"
	"io"
	"log"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"
)

const (
	csrftoken     = "csrftoken"
	README        = "README.md"
	LeetcodeQuery = "query question($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"
)

type Graphql struct {
	OperationName string `json:"operationName"`
	Variables     struct {
		TitleSlug string `json:"titleSlug"`
	} `json:"variables"`
	Query string `json:"query"`
}

var (
	lc     = &leetcodeCli{}
	logger = log.New(os.Stderr, "[leetcode-cli] ", log.Lshortfile|log.LstdFlags)
)

func init() {
	di.Put(lc)
	di.Put(simplehttp.NewClient(&simplehttp.Builder{ProtoMajor: simplehttp.HTTP2}))
}

type leetcodeCli struct {
	xCsrftoken string
	problems   *problems
	lock       sync.Mutex
	items      []item
	Commands   []cli.Command
	HttpClient simplehttp.Client `di:"*"`
	CmdConfig  Cmd               `di:"#.{cmd}"`
}

func (l *leetcodeCli) Init() {
	l.problems = new(problems)
}

func (l *leetcodeCli) AddCommand(command cli.Command) {
	l.Commands = append(l.Commands, command)
}

// readme添加item
func (l *leetcodeCli) AddItem(q *question) {
	var topics []string
	for e := range q.TopicTags {
		topics = append(topics, q.TopicTags[e].Name)
	}
	i := item{
		ID:         q.QuestionID,
		Title:      fmt.Sprintf("[%v](algorithms/%v/README.md)", q.Title, q.FileName),
		Topics:     strings.Join(topics, ", "),
		Difficulty: q.Difficulty,
	}
	for e := range l.items {
		if l.items[e].ID == q.QuestionID {
			l.items[e] = i
			return
		}
	}
	l.lock.Lock()
	l.items = append(l.items, i)
	l.lock.Unlock()
}

func (l *leetcodeCli) SortItem() {
	sort.Sort(itemSlice(l.items))
}

// 根据questionID生成 {question}.go README.md
func (l *leetcodeCli) Gen(ids []int) {
	p := pool.NewLimited(10)
	defer p.Close()
	batch := p.Batch()
	stats := l.problems.hasQuestionID(ids...)
	for e := range stats {
		stat := stats[e]
		batch.Queue(func(wu pool.WorkUnit) (r interface{}, err error) {
			descriptionUrl := fmt.Sprintf(l.CmdConfig.Params.Description, stat.QuestionTitleSlug)
			logger.Println("Fetch", stat.QuestionID, stat.QuestionTitle, descriptionUrl)
			err = simplehttp.NewRequest(l.HttpClient).Get().SetUrl(descriptionUrl).Send().Error()
			if err != nil {
				logger.Fatal(err)
				return
			}

			// 获取csrftoken，查询graphQL需要在header中添加x-csrftoken
			httpClient := l.HttpClient.(*simplehttp.HttpClient)
			parse, _ := url.Parse(l.CmdConfig.Params.LeetcodeCom)
			jar := httpClient.Jar.Cookies(parse)
			for _, v := range jar {
				if v.Name == csrftoken {
					l.xCsrftoken = v.Value
					break
				}
			}

			if l.xCsrftoken == "" {
				err = errors.New("x-csrftoken not in cookies")
				logger.Fatal(err)
				return
			}

			// 查询question
			qlResp, err := l.queryGraphQL(descriptionUrl, l.CmdConfig.Params.OperationName, stat.QuestionTitleSlug)
			if err != nil {
				logger.Fatal(err)
				return
			}
			q := newQuestion(&qlResp.Data.Question, l.CmdConfig.Generation.Algorithms)

			// 生成文件夹
			err = fs.MkdirP(q.DirPath)
			if err != nil {
				return
			}

			// 生成题解
			err = q.WriteSolution(l.CmdConfig.Template.Solution)
			if err != nil {
				return
			}

			err = q.WriteProblem(l.CmdConfig.Params.Lang, l.CmdConfig.Template.Problem)
			if err != nil {
				return
			}

			err = q.WriteUnittest(l.CmdConfig.Template.Unittest)
			if err != nil {
				return
			}

			l.AddItem(q)
			return
		})
	}
	batch.QueueComplete()
	batch.WaitAll()
}

// 读取algorithms.json
func (l *leetcodeCli) ReadLock() {
	err := fs.ReadJSON(l.CmdConfig.Lock.Algorithms, &l.items)
	if err != nil {
		logger.Println(err)
		panic(err)
	}
}

// 写algorithms.json
func (l *leetcodeCli) WriteLock() {
	err := fs.WriteJSON(l.CmdConfig.Lock.Algorithms, &l.items)
	if err != nil {
		logger.Println(err)
		panic(err)
	}
}

type graphQLResp struct {
	Data struct {
		Question problem `json:"question"`
	} `json:"data"`
}

// 查询题目
func (l *leetcodeCli) queryGraphQL(referer, operationName, titleSlug string) (*graphQLResp, error) {
	request := simplehttp.NewRequest(l.HttpClient)

	graphql := &Graphql{
		OperationName: operationName,
		Query:         LeetcodeQuery,
		Variables: struct {
			TitleSlug string `json:"titleSlug"`
		}{TitleSlug: titleSlug},
	}
	resp := request.Post().SetUrl(l.CmdConfig.Params.Graphql).
		Head("referer", referer).
		Head("x-csrftoken", l.xCsrftoken).
		Head("content-type", "application/json").
		SetJSON(graphql).
		Send()
	qlResponse := new(graphQLResp)
	err := resp.JSON(qlResponse)
	if err != nil {
		return nil, err
	}
	return qlResponse, nil
}

// 读取leetcode.json
func (l *leetcodeCli) ReadCatalog(path string, reGen bool) {
	if !fs.Exists(path) || reGen {
		l.WriteCatalog(path)
	}
	err := fs.ReadJSON(path, l.problems)
	if err != nil {
		panic(err)
	}
}

// 生成leetcode.json
func (l *leetcodeCli) WriteCatalog(path string) {
	_ = fs.WriteFileNotExist(path, func(writer io.Writer) error {
		b, err := simplehttp.NewRequest(l.HttpClient).Get().SetUrl(l.CmdConfig.Params.ProblemsAllURL).Send().Body()
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(b)
		if err != nil {
			panic(err)
		}
		return nil
	})
}

func (l *leetcodeCli) WriteReadme() {
	_ = fs.WriteFile(l.CmdConfig.Generation.Readme, func(writer io.Writer) error {
		tpl, err := template.ParseFiles(l.CmdConfig.Template.Readme)
		if err != nil {
			return err
		}
		err = tpl.Execute(writer, l.items)
		if err != nil {
			return err
		}
		return nil
	})
}
