package main

import "fmt"

func main() {
	name := "John cena"
	age := 44
	height := 6.2

	fmt.Printf("Name  :%s\n", name)             // %s is for strings
	fmt.Printf("Age : %d\n", age)               // %d is for integers
	fmt.Printf("Height :%.2f meters\n", height) // %.2f for float with 2 decimal places
}
