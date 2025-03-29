package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter a number:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	fmt.Println("Sum of first", n, "numbers is:", sum)
}
