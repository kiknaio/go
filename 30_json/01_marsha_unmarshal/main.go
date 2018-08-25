package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
