package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 测试中不建议这样写，原因：
	// 1.测试机器可能不具有外网访问的能力；
	// 2.有可能在测试时该网址挂掉了。
	// contents, err := fetcher.Fetch("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html") // 直接从本地读取，避免了测试对于外部环境的依赖

	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s\n", contents)
	res := ParseCityList(contents, "")
	expectedURLs := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}

	const resultSize = 470
	if len(res.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d",
			resultSize, len(res.Requests))
	}

	for i, url := range expectedURLs {
		if res.Requests[i].URL != url {
			t.Errorf("expected url #%d: %s, but had %s",
				i, url, res.Requests[i].URL)
		}
	}
}
