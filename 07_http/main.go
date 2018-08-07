package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func main()  {
	res, _ := http.Get("https://kikna.io")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	err := ioutil.WriteFile("index.html", page, 777)

	if err != nil {
		log.Fatal(err)
	}
}
