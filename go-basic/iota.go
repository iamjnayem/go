package main

import "fmt"


func main(){
	const (
		A = iota
		B 
		C
	)
	fmt.Println(A)


	const (
		_ = iota
		KB = 1 << (10)
		MB
		GB
	)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
