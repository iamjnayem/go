package main

import "fmt"

func main(){
	for i := 0; i < 1000000; i++ {
		go func(){
			fmt.Println(i)
		}()
	}
}
