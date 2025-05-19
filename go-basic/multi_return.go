package main 

import "fmt"


func myFunc(p, q int) (rectanngle int, square int){
	rectanngle = p * q 
	square = p * p 
	return 
}

func myTest(n1, n2 int) (int, int, int) {
	return n1 + n2, n1-n2, n1/n2

}

func main(){

	fmt.Println("Example of Multiple Value return from function")
	var rectanngle, square = myFunc(3, 2)
	fmt.Println("rectanngle", rectanngle)
	fmt.Println("square", square)

	add, _, _ := myTest(2, 3);
	fmt.Println("add",add)


}