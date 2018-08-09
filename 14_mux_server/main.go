package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"time"
	"log"
)

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}