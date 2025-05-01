package main

import (
"fmt"
"unsafe"
"strconv"
"sync"
)

func printAddr(xx int, wg *sync.WaitGroup){
	defer wg.Done()

	x := xx
	ptr := &x
	addr := uintptr(unsafe.Pointer(ptr))
	fmt.Println("=================================")
	fmt.Println("x = " + strconv.Itoa(x))
	fmt.Println("ptr = ", ptr)
	fmt.Println("addr = ", addr)
	fmt.Println("=================================")
	fmt.Println("")
	fmt.Println("")
}
func main(){

	var wg sync.WaitGroup

	for i:=0; i < 1000000000; i++ {
		go printAddr(i, &wg)
		wg.Add(1)
	}

	wg.Wait()
}

