package parser

import (
	"regexp"
	"strconv"

	"github.com/kentliuqiao/learngo/crawler/model"

	"github.com/kentliuqiao/learngo/crawler/engine"
)

var (
	// nameRe       = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	ageRe        = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hukouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xinzuoRe     = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

	guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://localhost:8080/mock/album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

	idURLRe = regexp.MustCompile(`http://localhost:8080/mock/album.zhenai.com/u/([\d]+)`)
)

// ParseProfile 用户信息解析器
func ParseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = extractString(contents, genderRe)
	// re := regexp.MustCompile(ageRe)
	// match := re.FindSubmatch(contents)
	// if match != nil {
	// 	age, err := strconv.Atoi(string(match[1]))
	// 	if err == nil {
	// 		profile.Age = age

	// 	}
	// }
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	// re = regexp.MustCompile(marriageRe)
	// match = re.FindSubmatch(contents)
	// if match != nil {
	// 	profile.Marriage = string(match[1])
	// }
	profile.Marriage = extractString(contents, marriageRe)

	res := engine.ParseResult{
		Items: []engine.Item{
			engine.Item{
				URL:     url,
				Type:    "zhenai",
				ID:      extractString([]byte(url), idURLRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		res.Requests = append(res.Requests, engine.Request{
			URL:        string(m[1]),
			ParserFunc: ProfileParser(string(m[2])),
		})
	}

	return res
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

// ProfileParser ...
func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
