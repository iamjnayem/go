package main 
import (
	"fmt"
	"sort"
)

func main(){
	intSlice := []int{42,23,16,15,8,4}
	fmt.Println("Before",intSlice)
	fmt.Println("Ints are sorted", sort.IntsAreSorted(intSlice))

	sort.Ints(intSlice)
	fmt.Println("After:",intSlice)
	fmt.Println("Ints are sorted", sort.IntsAreSorted(intSlice))

	sort.Slice(intSlice, func(i,j int) bool {
		return intSlice[i] > intSlice[j]
	})
	fmt.Println("Ints are now ", intSlice)

}