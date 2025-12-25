package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Define a struct to match the response structure
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// API URL to request
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching API data: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse the JSON response into a Post struct
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Output the parsed data
	fmt.Printf("Post ID: %d\n", post.ID)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}
