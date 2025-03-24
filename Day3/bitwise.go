package main

import "fmt"

func main() {
	a := 5
	b := 3

	fmt.Println("a & b:", a&b)
	fmt.Println("a | b:", a|b)
	fmt.Println("a ^ b:", a^b)
	fmt.Println("a << 1:", a<<1)
	fmt.Println("a >> 1:", a>>1)
}
