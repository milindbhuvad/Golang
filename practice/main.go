package main

import (
	"fmt"
	"net/http"

	// Import your local package
	reversestring "practice/reverse-string"
)

func main() {
	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		// Get the 'input' value from the URL query: /reverse?input=...
		input := r.URL.Query().Get("input")

		if input == "" {
			http.Error(w, "Missing 'input' parameter", http.StatusBadRequest)
			return
		}

		reversed := reversestring.Reverse(input)

		fmt.Fprintf(w, "Original: %s\nReversed: %s", input, reversed)
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
