package main

import "fmt"

func main() {
	fmt.Println("hello world")
	panic("idnia")
	fmt.Println("after panic")
    // unreachable code
    fmt.Println("this line will not be executed")
}

