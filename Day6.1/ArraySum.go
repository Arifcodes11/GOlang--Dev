package main

import "fmt"

func main(){
	arr:=[5]int{1,2,3,4,5}
	sum:=0

	for _, value:=range arr{
		sum=sum+value
	}
	fmt.Println("Sum of array elements:",sum)
}