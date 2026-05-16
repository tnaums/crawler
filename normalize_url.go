package main

import (
	"net/url"
	"fmt"
	"strings"
)

func normalizeURL(u string) (string, error) {
	p, err := url.Parse(u)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := p.Host + p.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
