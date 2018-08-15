package main

import "fmt"

// Closure
func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}