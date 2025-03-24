package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	result := num >> 2
	fmt.Println(num, "divided by 4 is : ", result)
}
