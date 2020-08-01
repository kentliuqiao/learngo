package persist

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/kentliuqiao/learngo/crawler/engine"
	"github.com/kentliuqiao/learngo/crawler/model"
	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
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

	// TODO: Try to start up elasticsearch
	// here using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.ID).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJSONObj(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v, expected %v", actual, expected)
	}
}
