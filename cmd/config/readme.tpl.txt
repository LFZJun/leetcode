<p align="center">
    <a href="https://github.com/ljun20160606/leetcode"><img src="doc/leetcode.jpeg" width="325"/></a>
</p>

<p align="center"> <code>Leetcode</code> 的 <code>Golang</code> 解法 😋</p>
<p align="center">
    🔥 <a href="#algorithms">Algorithms</a> |
    ✨ <a href="#other">Other</a>
</p>

<p align="center">
    <a href="https://github.com/ljun20160606/leetcode/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
    <a href="https://travis-ci.org/ljun20160606/leetcode"><img src="https://travis-ci.org/ljun20160606/leetcode.svg?branch=master"></a>
    <a href="https://codecov.io/gh/ljun20160606/leetcode"><img src="https://codecov.io/gh/ljun20160606/leetcode/branch/master/graph/badge.svg"></a>
    <a href="https://gitter.im/ljun20160606/leetcode?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge"><img src="https://badges.gitter.im/ljun20160606/leetcode.svg"></a>
    <a href="http://commitizen.github.io/cz-cli"><img src="https://img.shields.io/badge/commitizen-friendly-brightgreen.svg"></a>
</p>

***

## Algorithms

| #    | Title                                                                      | Topics                                            | Difficulty |
| :--- | :------------------------------------------------------------------------- | :------------------------------------------------ | :----------|
{{range .}}|{{.ID}} | {{.Title}} | {{.Topics}} | {{.Difficulty}} |
{{end}}
## Other

| #    | Title                                                  | Tag                 |
| :--- | :----------------------------------------------------- | :------------------ |
|      | [Ksum](algorithms/other/ksum.go)                       | Array, Two Pointers |
|      | [0/1 knapsack](algorithms/other/knapsack%20problem.go) | Dynamic Programming |

## Cmd

用于自动生成需要写的题目并重新生成README.md

```bash
# 进入cmd
$ cd cmd
# 获取依赖
$ go get -v -u
# 编译cli
$ ./build
# 返回项目目录
$ cd ..
# 拉题目
$ leetcode-cli pull 1
```
