package main

import "fmt"

func main() {
	country := "INDIA" // Go automatically infers type as string
	population := 1400 // Go infers type as int

	fmt.Println("Country:", country)
	fmt.Println("Population(in Millions):", population)
}
