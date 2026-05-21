package main

import (
	"fmt"
	"net/url"
)

// func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int)
func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()	

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

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

	if ! isFirst {
		return
	}
	fmt.Printf("Crawling %s...\n", rawCurrentURL)	
	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}
	//	fmt.Println(pageHTML)

	// Extract all the data we care about and store it
	pageData := extractPageData(pageHTML, rawCurrentURL)
	cfg.setPageData(normalized, pageData)
	

	for _, u := range pageData.OutgoingLinks {
		//		fmt.Println(u)
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	
	_, ok := cfg.pages[normalizedURL]
	if ok {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true
}

// setPageData safely stores the final PageData for a URL.
func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}
