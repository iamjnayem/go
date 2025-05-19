package main 

import "fmt"

func main() {
	a := []int{1,2}
	b := make([]int, 3)
	copy(b, a)
	b[2] = 3
	fmt.Println(b)
}