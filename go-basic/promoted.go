package main 

import "fmt"


type Address struct {}

func(a Address) showAddress(){
	fmt.Println("Showing address")
}



type Person struct {
	name string 
	Address
}


func main(){
	p := Person {
		name : "Liton",
	}
	p.showAddress()

	fmt.Println(p.name)
}