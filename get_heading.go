package main

import (
	//	"fmt"
	"strings"
	//"log"
	//"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	
	return doc.Find("h1").Text()
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	main := doc.Find("main")
	if main.Find("p").Text() != "" {
		return main.Find("p").Text()
	}
	return doc.Find("p").First().Text()
}
