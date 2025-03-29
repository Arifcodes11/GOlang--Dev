package main

import "fmt"

func main(){
	fmt.Println("Even numbers from 1 to 20")
	for i:=1;i<=20;i++{
		if i%2==0{
			fmt.Println(i)
		}
	}
}