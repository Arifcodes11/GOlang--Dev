package main

import "fmt"

func main() {
	var a, b int
	var tmp int
	fmt.Print("Enter two Numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("Before swap: a = ", a, " b = ", b)

	/*a = a ^ b
	b = a ^ b
	a = a ^ b*/
	tmp = a
	a = b
	b = tmp

	fmt.Println("After Swapping: a = ", a, "b = ", b)
}
