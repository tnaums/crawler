package main

import (
	"fmt"
	"os"
)

var pages = make(map[string]int)

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
	fmt.Printf("starting crawl of: %s\n", baseURL)
	
	crawlPage(baseURL, baseURL, pages)

	fmt.Println("Done. Printing map:")
	for k, v := range pages {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}
