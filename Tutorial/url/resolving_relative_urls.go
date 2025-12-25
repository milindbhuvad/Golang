package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL, _ := url.Parse("https://www.example.com/")
	relURL, _ := url.Parse("path/to/resource")

	fullURL := baseURL.ResolveReference(relURL)
	fmt.Println("Resolved URL:", fullURL.String())
}
