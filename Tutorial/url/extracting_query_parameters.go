package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://www.example.com/search?query=golang&sort=asc"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Extract query parameters
	queryParams := parsedURL.Query()
	fmt.Println("Query parameter 'query':", queryParams.Get("query"))
	fmt.Println("Query parameter 'sort':", queryParams.Get("sort"))
}
