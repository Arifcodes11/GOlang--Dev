// Slices share the same underlying array, so modifying one slice may affect the original or other slices derived from it.
package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice1 := arr[1:4] // Slice of elements [20, 30, 40]

	fmt.Println("Original Array:", arr)
	fmt.Println("Original Slice1:", slice1)

	// Modifying the slice
	slice1[1] = 99
	fmt.Println("Modified Slice1:", slice1)
	fmt.Println("Array After Modification:", arr)
}
