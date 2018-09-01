package main

import (
"fmt"
	"runtime"
	"sync"
"time"
)

var wg sync.WaitGroup

func init() {
	// Use all the CPUs for goroutines
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 45; i++  {
		fmt.Println("Foo:🥕", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++  {
		fmt.Println("Bar:🍗", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}