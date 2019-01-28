package main

import "strings"

type Language struct {
	Lang      string
	Extension string
}

var (
	Cpp = Language{
		Lang:      "C++",
		Extension: "cpp",
	}
	C = Language{
		Lang:      "C",
		Extension: "c",
	}
	Java = Language{
		Lang:      "Java",
		Extension: "java",
	}
	Python = Language{
		Lang:      "Python",
		Extension: "py",
	}
	Python3 = Language{
		Lang:      "Python3",
		Extension: "py",
	}
	JavaScript = Language{
		Lang:      "JavaScript",
		Extension: "js",
	}
	Go = Language{
		Lang:      "Go",
		Extension: "go",
	}

	langs = []Language{
		Cpp, C, Java, Python, Python3, JavaScript, Go,
	}
)

func MatchLang(str string) *Language {
	for i := range langs {
		language := langs[i]
		if strings.EqualFold(language.Lang, str) {
			return &language
		}
	}
	return nil
}
