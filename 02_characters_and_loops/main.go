package main

import "fmt"

func main() {
	//Print numbers in decimal, binary, hexidecimal and utf-8
	for i:= 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t 0x%x - %q\n", i, i, i, i)
	}
}
