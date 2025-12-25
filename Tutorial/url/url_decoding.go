package main

import (
	"fmt"
	"net/url"
)

func main() {
	encodedString := "Golang+is+great%21+%231"
	decoded, err := url.QueryUnescape(encodedString)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded string:", decoded)
}
