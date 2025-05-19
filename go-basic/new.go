package main

import "fmt"

func main(){
	p := new(int)
	fmt.Println(p)

	*p = 420
	fmt.Println(*p)

	
}