package main

import "fmt"

func main(){
	c := complex(1.0, 2.0);
	fmt.Println(c)

	fmt.Println(real(c))
	fmt.Println(imag(c))
	
}