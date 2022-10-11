package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	Seed(100)
	r := Random()
	fmt.Printf("r = %+v\n", r)
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
