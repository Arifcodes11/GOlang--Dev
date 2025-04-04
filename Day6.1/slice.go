package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	fmt.Println("Origibal slicd:", nums)

	nums = append(nums, 40, 50, 60)
	fmt.Println("Updated slice after append:", nums)
}
