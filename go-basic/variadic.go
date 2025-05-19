package main 

import "fmt"


func myFunc(nums ...int){
	for x, num := range nums {
		fmt.Println(x,num)
	}
}

func sum(nums ...int){
	total := 0 

	for _, num := range nums {
		total += num 
	}

	fmt.Println("Sum: ", total)
}

func main() {
	sum(1, 2, 3)
}