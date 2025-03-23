package main

import "fmt"

func main() {
	message := "Go is Awesome!" // Automatically inferred as string
	number := 100               // Inferred as int

	fmt.Println(message, number)
}
