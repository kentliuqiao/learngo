package parser

import (
	"regexp"

	"github.com/kentliuqiao/learngo/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURLRe = regexp.MustCompile(`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)
)

// ParseCity 城市页解析器
func ParseCity(contents []byte, _ string) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	res := engine.ParseResult{}
	for _, m := range matches {
		// res.Items = append(res.Items, "User "+name) // 去除无用数据
		res.Requests = append(res.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}

	matches = cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		res.Requests = append(res.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return res
}
