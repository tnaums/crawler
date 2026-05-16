package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	extractedImages := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return extractedImages, err
	}

	images := doc.Find("img")
	photograph, _ := images.Attr("src")
	fmt.Printf("Type of images: %T\n", images)
	fmt.Printf("Type of photograph: %T\n", photograph)
	fmt.Println(photograph)
	fmt.Println("Yo!")

	return extractedImages, nil
}
