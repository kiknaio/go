package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the WaitGroup
var wg sync.WaitGroup

func foo() {
	for i := 0; i < 45; i++  {
		fmt.Println("Foo:🥕", i)
		// Wait for 3 milliseconds and then run
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	// When loop is finished remove one item from WaitGroup
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++  {
		fmt.Println("Bar:🍗", i)
		// Wait for 10 milliseconds and then run
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	// When loop is finished remove one item from WaitGroup
	wg.Done()
}

func main() {
	// Add two items in WaitGroup
	wg.Add(2)
	// These two processes wil run concurrently
	go foo()
	go bar()
	// Exit this scope when the WaitGroup will be empty
	wg.Wait()
}
