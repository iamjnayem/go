package main

import "fmt"

type Animal interface{
	Speak() string
}

type Canine struct{}

func(c Canine) Bark() string{
	return "woof"
}

type Dog struct{
	Canine // embedding
	Name string
}

func main(){
	d := Dog{Name:"kira"}
	fmt.Println(d.Bark())
	fmt.Println(d.Speak())
}
