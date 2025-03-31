package main
import "fmt"

func main(){
	var numbers [5]int

	numbers[0]=1
	numbers[1]=2
	numbers[2]=3
	numbers[3]=4
	numbers[4]=5

	fmt.Println("Array elements:")
	for i:=0;i<5;i++{
		fmt.Println(numbers[i])
	}
}