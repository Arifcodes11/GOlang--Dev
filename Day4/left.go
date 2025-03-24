package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a Number : ")
	fmt.Scan(&num)

	result := num << 1
	fmt.Println(num, " multiplied byu 2 is : ", result)
}
