package main
import "fmt"

var source = [5]int{10, 20, 30, 40, 50}

func main() {
    // Creating a pointer to the source array
    // var destination *[5]int = &source
    destination := &source

    fmt.Println("Source:", source)
    fmt.Println("Destination Array via pointer:", *destination)
}