package main 

import "fmt"

func process(callback func(int)){
	for i := 1; i <= 3; i++{
		callback(i)
	}
}

func mulitiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}

}

func counter() func() int {
	count := 0 
	return func() int {
		count++ 
		return count 
	}
}

func main(){
	name := "Liton"
	sayHello := func(){
		fmt.Println("Hello",name)
	}

	sayHello()  

	myCallback := func(n int){
		fmt.Println("I am called with",n)
	}

	process(myCallback)

	times2 := mulitiplier(2)
	times3 := mulitiplier(3)

	fmt.Println(times2(5))
	fmt.Println(times3(5))

	c := counter()

	fmt.Println(c())
	fmt.Println(c())
}