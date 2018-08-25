package main

import (
	"net/http"
	"log"
	"fmt"
	"bufio"
)

func HashBucket(word string) int {
	return int(word[0])
}

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")

	if err != nil {
		log.Fatalln(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:122])
}
