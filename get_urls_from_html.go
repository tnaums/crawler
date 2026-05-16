package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	extractedURLs := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return extractedURLs, err
	}

	u := doc.Find("a").Text()
	fmt.Println(u)
	fmt.Println("Yo!")

	return extractedURLs, nil
}
