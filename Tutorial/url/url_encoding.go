package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawString := "Golang is great! #1"
	escaped := url.QueryEscape(rawString)
	fmt.Println("Escaped string:", escaped)
}
