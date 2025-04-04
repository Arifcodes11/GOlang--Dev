package main

import "fmt"

func main(){
	nums := []int{15,22,8,43,19}
	max:=nums[0]

	for _, value := range nums{
		if value >max{
			max=value
		}
	}
	fmt.Println("Maxiumum value in the slice:",max)
}
