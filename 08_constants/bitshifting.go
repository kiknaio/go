package main

import "fmt"

// Bit shifting (Bitwise)
// Shift 1 to left by 10
const (
	D = 1 << 10 // 1024
)

func main() {
	fmt.Printf("%d", D)
}
