package engine

import (
	"log"

	"github.com/kentliuqiao/learngo/crawler/fetcher"
)

func worker(r Request) (ParseResult, error) {
	// log.Printf("fetching url: %s", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("fetcher error: %v, url: %s", err, r.URL)
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.URL), nil
}
