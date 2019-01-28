package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/ljun20160606/gox/fs"
	"github.com/pkg/errors"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

type question struct {
	*problem
	FileName string
	DirPath  string
}

type solutionTemplate struct {
	QuestionTitle string
	Content       string
}

// 在指定目录下生成题目
func newQuestion(problem *problem, dirPrefix string) *question {
	fileName := offsetNumber(problem.QuestionID, '0', 3)
	return &question{
		problem:  problem,
		FileName: fileName,
		DirPath:  filepath.Join(dirPrefix, fileName),
	}
}

// 生成题解
func (q *question) WriteSolution(solutionTplPath string) error {
	return fs.WriteFileNotExist(filepath.Join(q.DirPath, README), func(writer io.Writer) error {
		tpl, err := template.ParseFiles(solutionTplPath)
		if err != nil {
			return err
		}
		document, err := goquery.NewDocumentFromReader(strings.NewReader(q.Content))
		if err != nil {
			return err
		}
		s := solutionTemplate{
			QuestionTitle: q.Title,
			Content:       document.Text(),
		}
		err = tpl.Execute(writer, s)
		if err != nil {
			return err
		}
		return nil
	})
}

// 生成题目 problem
func (q *question) WriteProblem(lang, problemTpl string) error {
	language := MatchLang(lang)
	if language == nil {
		return errors.Errorf("local not support %v.", language)
	}
	return fs.WriteFileNotExist(filepath.Join(q.DirPath, q.QuestionFrontendID+". "+q.Title+"."+language.Extension), func(writer io.Writer) error {
		tpl, err := template.ParseFiles(problemTpl)
		if err != nil {
			return err
		}
		snippet := getCodeSnippet(q.CodeSnippets, *language)
		if snippet == nil {
			return errors.New("remote not support " + Go.Lang + ".")
		}
		err = tpl.Execute(writer, snippet)
		if err != nil {
			return err
		}
		return nil
	})
}

func getCodeSnippet(cs []codeSnippet, expect Language) *codeSnippet {
	for j := range cs {
		if cs[j].Lang == expect.Lang {
			return &cs[j]
		}
	}
	return nil
}

// 生成题目 problem 的单元测试模版
func (q *question) WriteUnittest(unittestTpl string) error {
	return fs.WriteFileNotExist(filepath.Join(q.DirPath, q.QuestionFrontendID+"_test.go"), func(writer io.Writer) error {
		tpl, err := template.ParseFiles(unittestTpl)
		if err != nil {
			return err
		}
		err = tpl.Execute(writer, map[string]interface{}{"ID": q.QuestionFrontendID})
		if err != nil {
			return err
		}
		return nil
	})
}
