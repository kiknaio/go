package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	// Convert string to slice of bytes
	bs := []byte(`{"First": "James", "Last": "Bond", "Age": 20}`)
	json.Unmarshal(bs, &p1) // Point to a address where to save

	fmt.Println("-----------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}