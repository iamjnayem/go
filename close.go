package main

import "fmt"


func main(){
	ch := make(chan int)
	close(ch)
}