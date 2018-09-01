// Channels are like marathon runners
package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel is similar to maps
	// Unbuffered channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++  {
			// Putting number into channel
			// Here goroutine stops until something
			// takes value from here
			c <- i
		}
	}()

	go func() {
		for {
			// Retrieving number from the channel
			fmt.Println(<-c)
		}
	}()
	// synthetically wait for 1 second to run all the above processes
	time.Sleep(time.Second)
}
