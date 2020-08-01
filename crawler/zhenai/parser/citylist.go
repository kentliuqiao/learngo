package parser

import (
	"regexp"

	"github.com/kentliuqiao/learngo/crawler/engine"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList 解析城市列表页
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	res := engine.ParseResult{}
	for _, m := range matches {
		// res.Items = append(res.Items, "City "+string(m[2])) // 去除无用数据
		res.Requests = append(res.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return res
}
