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
	if err != nil {
		return
	}

	isFirst := cfg.addPageVisit(normalized)

	if isFirst == false {
		return
	}

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

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	_, ok := cfg.pages[normalizedURL]
	if ok {
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1

	return true
}
