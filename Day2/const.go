package main

import "fmt"

func main() {
	const PI float64 = 3.14159    //Explicit type
	const Greeting = "Hello, Go!" //implicit type

	fmt.Println("PI:", PI)
	fmt.Println(Greeting)
}
