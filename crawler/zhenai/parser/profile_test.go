package parser

import (
	"io/ioutil"
	"testing"

	"github.com/kentliuqiao/learngo/crawler/engine"
	"github.com/kentliuqiao/learngo/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("./profile_test_data.html")
	if err != nil {
		panic(err)
	}

	res := ParseProfile(contents,
		"http://localhost:8080/mock/album.zhenai.com/u/4143921402121684230", "何必怀念深碍")
	if len(res.Items) != 1 {
		t.Errorf("result items should contain 1 element, but was %v", res.Items)
	}

	actual := res.Items[0]

	expected := engine.Item{
		URL:  "http://localhost:8080/mock/album.zhenai.com/u/4143921402121684230",
		Type: "zhenai",
		ID:   "4143921402121684230",
		Payload: model.Profile{
			Name:       "何必怀念深碍",
			Gender:     "女",
			Age:        93,
			Height:     56,
			Weight:     281,
			Income:     "财务自由",
			Marriage:   "离异",
			Education:  "初中",
			Occupation: "销售",
			Hukou:      "天津市",
			House:      "租房",
			Car:        "有车",
			Xinzuo:     "白羊座",
		},
	}
	if actual != expected {
		t.Errorf("expected %v, but was %v", expected, actual)
	}
}
