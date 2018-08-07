package main

import "fmt"

const (
	name string = "George"
	age = 23
	A = iota
	B = iota
	C = iota
)

func main() {
	println(name)
	println(age)


	fmt.Printf("%d %d %d\n", A, B, C)
}
