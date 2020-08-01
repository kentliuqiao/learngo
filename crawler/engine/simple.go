package engine

import (
	"log"
)

// SimpleEngine 单任务版爬虫
type SimpleEngine struct{}

// Run drives the crawler
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseRes, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseRes.Requests...)
		for _, item := range parseRes.Items {
			log.Printf("got item: %v\n", item)
		}
	}
}
