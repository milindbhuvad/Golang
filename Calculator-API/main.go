package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CalcResponse defines the structure of our JSON answer
type CalcResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Get values from URL query: ?a=10&b=5&op=add
	valA := r.URL.Query().Get("a")
	valB := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")

	// 2. Convert strings to numbers (float64)
	a, _ := strconv.ParseFloat(valA, 64)
	b, _ := strconv.ParseFloat(valB, 64)

	var result float64
	var errStr string

	// 3. Perform the math logic
	switch op {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			errStr = "cannot divide by zero"
		} else {
			result = a / b
		}
	default:
		errStr = "invalid operator (use add, sub, mul, or div)"
	}

	// 4. Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CalcResponse{
		Result: result,
		Error:  errStr,
	})
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)

	fmt.Println("Calculator Service running on :8081...")
	http.ListenAndServe(":8081", nil)
}
