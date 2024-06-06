package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	uiServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", uiServer)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
