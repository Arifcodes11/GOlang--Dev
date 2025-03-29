package main

import "fmt"

func main() {
	nums := [5]int{1, 3, 7, 8, 9}

	for index, value:=range nums {
		fmt.Println("index:", index, " value:", value)
	}
}
