package main 

import (
	"bytes"
	"fmt"
)

func main(){
	data := []byte("a,b,c")
	parts := bytes.Split(data, []byte(","))
	fmt.Println("Split Parts")
	for _, part := range parts {
		fmt.Println(string(part))
	}
}