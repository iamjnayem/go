package main

import "fmt"

func main(){
	m := map[string]int{"a":1}

	delete(m, "a")
	fmt.Println(m)

	m["b"] = 3

	fmt.Println(m)
}