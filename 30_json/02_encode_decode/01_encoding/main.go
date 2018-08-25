package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	// Stream Person to encoder and get JSON as a result
	json.NewEncoder(os.Stdout).Encode(p1)
}
