package engine

import (
	"learngo.com/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fecthing url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// 把获得的内容交由Parser来解析
	return r.ParserFunc(body), nil
}