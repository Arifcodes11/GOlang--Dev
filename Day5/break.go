package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}
}
