package main

import "fmt"

func main(){
	// nums := []int{10, 20, 30}
	// for i, v := range nums {
	// 	fmt.Println("index: ", i)
	// 	fmt.Println("v: ", v)
	// }

	// countries := []string {"Bangladesh", "India", "Pakistan"}
	// for i,j := range countries{
	// 	fmt.Println(i,j)
	// }

	// homeWork := "আমি আমার বাবাকে ভালোবাসি"

	// for start := 0; start < 1000000; start++ {
	// 	fmt.Println(homeWork)
	// }

	for i, j := range "abcd" {
		fmt.Printf("%U %d\n", j, i)
	}

	myMap := map[int]string {
		22: "liton",
		30: "nayem",
		44: "pulok",
	}

	for key, value := range myMap {
		fmt.Println(key,value)
	}

	
}