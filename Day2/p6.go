package main

import "fmt"

func main() {
	var a int = 10
	var b float64 = 5.5

	var sum float64 = float64(a) + b // Convert int to float64 before adding
	var sumInt int = int(b) + a      // Convert float64 to int before adding

	fmt.Println("Sum a float: ", sum)
	fmt.Println("Sum ad int: ", sumInt)
}
