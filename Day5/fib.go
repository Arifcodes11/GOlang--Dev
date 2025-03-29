package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)

	a, b := 0, 1 // First two terms

	fmt.Println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		fmt.Println(a)
		a, b = b, a+b // Update a and b for next iteration
	}
}
