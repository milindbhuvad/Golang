package main

import "fmt"

func main() {
	var x int = 5 // Declaring a variable x with type int
	var y = 10    // Go infers the type (int in this case)
	z := 15       // Short form declaration, infers the type as well

	fmt.Println(x, y, z)
}
