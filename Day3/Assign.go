package main

import "fmt"

func main() {
	x := 5
	fmt.Println("Initial value:", x) // x = 5

	x += 3                        // x = x + 3
	fmt.Println("After += 3:", x) // x = 8

	x -= 2                        // x = x - 2
	fmt.Println("After -= 2:", x) // x = 6

	x *= 4                        // x = x * 4
	fmt.Println("After *= 4:", x) // x = 24

	x /= 6                        // x = x / 6
	fmt.Println("After /= 6:", x) // x = 4

	x %= 3                        // x = x % 3
	fmt.Println("After %= 3:", x) // x = 1
}
