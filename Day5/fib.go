package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms : ")
	fmt.Scan(&n)

	a,b:=0,1

	fmt.Println("fibonacci series:")
	for i:=0;i<n;i++{
		fmt.Println(a)
		a,b=b,a+b
	}
}
