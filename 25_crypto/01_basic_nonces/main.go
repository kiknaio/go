package main

import (
	"fmt"
	"io"
	"crypto/rand"
)

func main() {
	var nonce [24]byte

	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
}
