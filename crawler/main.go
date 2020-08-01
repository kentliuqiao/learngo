package main

import (
	"github.com/kentliuqiao/learngo/crawler/engine"
	"github.com/kentliuqiao/learngo/crawler/persist"
	"github.com/kentliuqiao/learngo/crawler/scheduler"
	"github.com/kentliuqiao/learngo/crawler/zhenai/parser"
)

func main() {
	// engine.SimpleEngine{}.Run(engine.Request{
	// 	URL:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })
	// e := engine.ConcurrentEngine{
	// 	Scheduler:   &scheduler.SimpleScheduler{},
	// 	WorkerCount: 10,
	// }
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		URL:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	// e.Run(engine.Request{
	// 	URL:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
