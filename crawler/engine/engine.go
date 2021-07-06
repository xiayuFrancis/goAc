package engine

import (
	"goAc/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error  fetching url %s : %v", r.Url, err)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("got item :%v", item)
		}
	}
}
