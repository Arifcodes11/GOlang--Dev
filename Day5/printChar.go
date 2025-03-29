package main

import "fmt"

func main() {
	str := "Hello Go!"
	fmt.Println("Characters in the range :")

	for index, char := range str {
		fmt.Println("index:%d,character: %c\n",index,  char)
	}
}
