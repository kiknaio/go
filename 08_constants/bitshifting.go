package main

import "fmt"

// Bit shifting (Bitwise)
// Shift 1 to left by 10
const (
	_ = iota // 0
	KB = 1 << (iota * 10) // 1024 bytes
	MB = 1 << (iota * 10) // 1048576 bytes
	GB = 1 << (iota * 10) // 1073741824 bytes
)

func main() {
	fmt.Printf("KB: %d - %b\n", KB, KB)
	fmt.Printf("MB: %d - %b\n", MB, MB)
	fmt.Printf("GB: %d - %b\n", GB, GB)

	fmt.Println("willy")
}
