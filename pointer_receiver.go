package main 

import "fmt"

type Person struct {
	name string 
}

func(p *Person) changeName(newName string){
	p.name = newName
}

func(p Person) showName(){
	fmt.Println("Name: ", p.name)
}


func main(){
	p1 := Person{name: "Liton"}
	p1.changeName("Pulok")
	fmt.Println("Updated Name",p1.name)

	amir := Person{name: "Amir"}
	amir.changeName("Kamal")
	(&amir).showName()



}