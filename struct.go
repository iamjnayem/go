package main 

import "fmt"

type Address struct{
	city string 
	zip string 
}

type Person struct{
	name string 
	age int 
	address Address 
}

func main(){
	p := Person {
		name: "Liton",
		age: 23,
		address : Address {
			city: "Dhaka",
			zip : "1320",
		},
	}

	fmt.Println(p.name)
	fmt.Println(p.address)

	p1 := struct {
		name string 
		age int 
	}{
		name: "Hasina",
		age : 86,
	}

	fmt.Println(p1)


}