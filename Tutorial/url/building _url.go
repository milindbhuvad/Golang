package main

import (
	"fmt"
	"net/url"
)

func main() {
	base := &url.URL{
		Scheme: "https",
		Host:   "www.example.com",
		Path:   "/search",
	}

	// Add query parameters
	query := url.Values{}
	query.Add("query", "golang")
	base.RawQuery = query.Encode()

	// Output the full URL
	fmt.Println("Full URL:", base.String())
}
