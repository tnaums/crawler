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
	maxPages           int
}

func NewConfig(baseURL string, maxConcurrency int, maxPages int) config {
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
		concurrencyControl: make(chan struct{}, maxPages),
		wg:                 &sync,
		maxPages:           maxPages,
	}
}

func main() {
	var maxConcurrency int
	var maxPages int

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <url> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	if _, err := fmt.Sscanf(os.Args[2], "%d", &maxConcurrency); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	if _, err := fmt.Sscanf(os.Args[3], "%d", &maxPages); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Printf("type of maxConcurrency: %T\n", maxConcurrency)
	fmt.Printf("type of maxPages: %T\n", maxPages)
	crawlConfig := NewConfig(baseURL, maxConcurrency, maxPages)

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
