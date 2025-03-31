//2. Slices in Go
//Slices are dynamic, flexible, and more powerful than arrays. They are built on top of arrays and can grow or shrink.
//var slice []type
//Uses square brackets without specifying the size.

//Declaring and Using a Slice

package main
import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}  // Slice initialization
    fmt.Println("Original Slice:", nums)

    // Appending elements
    nums = append(nums, 6, 7)
    fmt.Println("After Append:", nums)

    // Slicing
    part := nums[1:4]  // From index 1 to 3 (exclusive of 4)
    fmt.Println("Sliced Part:", part)
}
