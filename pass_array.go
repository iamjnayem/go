package main 

import "fmt"


func calculateAverage(arr [6]int, size int) float32 {

	sum := 0 

	for _, value := range(arr) {
		sum += value 
	}


	return float32(sum) / float32(size) 
}


func calculateAveragePointerVersion(arr *[6]int, size int) float32{
	sum := 0 
	for _, value := range(arr){
		sum += value 
	}

	defer func(){
		arr[5] = 10
	}()

	return float32(sum)/float32(size)
}

func incrementEachItem(arr *[6]int){
	for i := range(arr){
		arr[i]++
	}
}




func main() {
	myArray := [6]int {1, 2, 3, 4, 5, 6}

	result := calculateAverage(myArray, 6)
	fmt.Println(result)

	anotherResult := calculateAveragePointerVersion(&myArray, 6)
	fmt.Println(anotherResult)
	fmt.Println("last index",myArray[5])

	incrementEachItem(&myArray)
	fmt.Println(myArray)
}