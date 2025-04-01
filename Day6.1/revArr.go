package main
import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}  // Original array
    n := len(arr)

    // Reversing the array
    for i := 0; i < n/2; i++ {
        arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
    }

    fmt.Println("Reversed array:", arr)
}
