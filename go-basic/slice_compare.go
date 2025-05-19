package main 

import (
	"fmt"
	"bytes"
	"reflect"
)

func main(){
	slice1 := []byte {'I', 'a','m', ' ', 'g', 'o', 'o', 'd', ' ', 'b', 'o','y'}
	slice2 := []byte {'I', 'a','m', ' ', 'g', 'o', 'o', 'd', ' ', 'b', 'o','y'}	
	slice3 := []byte {'I', 'a','m', ' ', 'g', 'o', 'o', 'd', ' ', 'g', 'i','r','l'}

	fmt.Println(bytes.Compare(slice1,slice2))
	fmt.Println(bytes.Compare(slice1,slice3))

	fmt.Println(bytes.Equal(slice1, slice2))
	fmt.Println(bytes.Equal(slice2, slice3))


	fmt.Println(reflect.DeepEqual(slice1,slice2))
	fmt.Println(reflect.DeepEqual(slice1,slice3))


}
