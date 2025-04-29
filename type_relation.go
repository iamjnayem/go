package main

import "fmt"
import "strconv"


type Animal interface{
	Speak() string
}


type Dog struct{
	Name string
}

type Cat struct{
	Name string
	Age int 
}


func(d Dog) Speak() string {
	return d.Name + "says woff"
}

func(c Cat) Speak() string{
	return c.Name + " is crying since " + strconv.Itoa(c.Age) + "days"
}


func main(){
	dog := Dog{ Name:"Charley"}
	cat := Cat{ Name:"Kitty", Age:1}
	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())
}
