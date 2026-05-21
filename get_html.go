package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func getHTML(rawURL string) (string, error) {
	scraperClient := NewClient(5 * time.Second)

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("error initializing request: %v", err)
	}

	req.Header.Set("User-Agent", "BootCrawler/1.0")
	resp, err := scraperClient.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v\n", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	//fmt.Println("Response Status Code:", resp.StatusCode)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("response status code %d", resp.StatusCode)
	}

	// Check if response is html
	ct := resp.Header.Get("content-type")
	if strings.Contains(ct, "text/html") != true {
		return "", fmt.Errorf("response is not text/html")
	}

	// read response body into []byte
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)

	}
	
	return string(data), nil
}
