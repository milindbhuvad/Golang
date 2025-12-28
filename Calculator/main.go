package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	result := calculate(num1, num2, operator)
	fmt.Println("Result:", result)
}

func calculate(a float64, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Error: Division by zero")
			return 0
		}
		return a / b
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}
