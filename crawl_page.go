package main

import (
	"fmt"
	"net/url"
)

// func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int)
func (cfg *config) crawlPage(rawCurrentURL string) {
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil { return }

	_, ok := cfg.pages[normalized]
	if ok {
		cfg.pages[normalized]++
		return
	}
	cfg.pages[normalized] = 1

	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}
	fmt.Println(pageHTML)

	allURLs, err := getURLsFromHTML(pageHTML, cfg.baseURL)
	if err != nil {
		return
	}
	fmt.Println("Looping...")
	for _, u := range allURLs {
		fmt.Println(u)
		cfg.crawlPage(u)
	}


}
