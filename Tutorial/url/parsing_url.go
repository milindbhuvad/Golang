package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://www.example.com:8080/path?query=golang#fragment"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Print the parsed URL components
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
}
