package main

type Cmd struct {
	Lock struct {
		Leetcode   string `yaml:"leetcode"`
		Algorithms string `yaml:"algorithms"`
	} `yaml:"lock"`
	Generation struct {
		Readme     string `yaml:"readme"`
		Algorithms string `yaml:"algorithms"`
	} `yaml:"generation"`
	Params struct {
		OperationName  string `yaml:"operation-name"`
		LeetcodeCom    string `yaml:"leetcode-com"`
		Graphql        string `yaml:"graphql"`
		ProblemsAllURL string `yaml:"problems-all-url"`
		Description    string `yaml:"description"`
		Lang           string `yaml:"lang"`
	} `yaml:"params"`
	Template struct {
		Solution string `yaml:"solution"`
		Problem  string `yaml:"problem"`
		Readme   string `yaml:"readme"`
		Unittest string `yaml:"unittest"`
	} `yaml:"template"`
}
