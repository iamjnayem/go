package main 

import (
	"fmt"
	"reflect"
)

type Author struct {
	name string
	branch string 
	language string 
	particles int
}

func main(){
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s\n", rune1, rune1, reflect.TypeOf(rune1))
	fmt.Printf("Rune 2: %c; Unicode: %U; Type: %s\n", rune2, rune2, reflect.TypeOf(rune2))
	fmt.Printf("Rune 3: %c; Unicode: %U; Type: %s\n", rune3, rune3, reflect.TypeOf(rune3))

	a1 := Author {
		name: "Meghna",
		branch : "Dhaka",
		language : "Go",
		particles : 38,
	}

	a2 := Author {
		name: "Meghna",
		branch: "Dhaka",
		language: "Go",
		particles: 38,
	}

	a3 := Author{
		name: "Kona",
		branch: "Barishal",
		language: "Java",
		particles: 39,

	}

	if a1 == a2 {
		fmt.Println("both teacher are is same");
	}

	fmt.Println(a3)

	fmt.Println("a1 & a2 are same: ", reflect.DeepEqual(a1, a2))

}