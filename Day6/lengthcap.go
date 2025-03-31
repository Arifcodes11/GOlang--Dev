// Length and Capacity
// Length (len()): Number of elements in the slice.

// Capacity (cap()): Number of elements in the underlying array from the starting index of the slice.

package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	slice1 := nums[1:4] // Slice of elements [20, 30, 40]

	fmt.Println("Lenght of Slice1:", len(slice1))   // 3
	fmt.Println("Capacity of Slice1:", cap(slice1)) // 4
}
