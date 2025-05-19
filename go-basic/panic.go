package main

import "fmt"


func main(){
	fmt.Println("something is working")
	panic(1/0)
	fmt.Println("Not reachable code")
}