package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num&1 == 0 {
		fmt.Print(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
