package main 

import "fmt"
 

func main(){
	var x float32 = 5;
	var y float32 = 2.25;

	add := x + y 
	sub := x - y 
	div := x / y 
	mul := x * y 


	fmt.Printf("%g + %g = %f\n", x,y,add)
	fmt.Printf("%f - %f = %g\n", x,y,sub)
	fmt.Printf("%f * %f = %f\n", x,y,mul)
	fmt.Printf("%f / %f = %f\n", x,y,div)



}