package main

import (
	"encoding/json"
	"net/http"
)

type CalculationRequest struct {
	Operation string  `json:"operation"`
	Operand1  float64 `json:"operand1"`
	Operand2  float64 `json:"operand2"`
}
type CalculationResponse struct {
	Result float64 `json:"result"`
}

func calculate(w http.ResponseWriter, r *http.Request) {
	var req CalculationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var result float64
	switch req.Operation {
	case "add":
		result = req.Operand1 + req.Operand2
	case "subtract":
		result = req.Operand1 - req.Operand2
	case "multiply":
		result = req.Operand1 * req.Operand2
	case "divide":
		if req.Operand2 == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		}
		result = req.Operand1 / req.Operand2
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}
	resp := CalculationResponse{Result: result}
	json.NewEncoder(w).Encode(resp)
}
func main() {
	http.HandleFunc("/calculate", calculate)
	http.ListenAndServe(":8080", nil)
}
