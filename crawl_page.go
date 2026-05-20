package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil { return }

	_, ok := pages[normalized]
	if ok {
		pages[normalized]++
		return
	}
	pages[normalized] = 1

	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}
	fmt.Println(pageHTML)

	allURLs, err := getURLsFromHTML(pageHTML, parsedBaseURL)
	if err != nil {
		return
	}
	fmt.Println("Looping...")
	for _, u := range allURLs {
		fmt.Println(u)
		crawlPage(rawBaseURL, u, pages)
	}


}
