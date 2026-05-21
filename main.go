package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

//var pages = make(map[string]int)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func NewConfig(baseURL string) config {
	pages := make(map[string]PageData)

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return config{}
	}

	var mutex sync.Mutex
	var sync sync.WaitGroup

	return config{
		pages:              pages,
		baseURL:            parsedBaseURL,
		mu:                 &mutex,
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync,
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]

	crawlConfig := NewConfig(baseURL)

	fmt.Printf("starting crawl of: %s\n", baseURL)

	crawlConfig.wg.Add(1)
	go crawlConfig.crawlPage(baseURL)

	fmt.Println("Done. Printing map:")
	for k, v := range crawlConfig.pages {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	crawlConfig.wg.Wait()

	for k, _ := range crawlConfig.pages {
		fmt.Printf("found: %s\n", k)
	}
}
