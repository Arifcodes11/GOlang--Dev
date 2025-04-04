package main
import "fmt"

func main(){
	matrix:=[][]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	sum:=0
	for _, row:=range matrix{
		for _, value := range row{
			sum+=value
		}
	}
	fmt.Println("Sum of all elements in the 2d slice:",sum)
}