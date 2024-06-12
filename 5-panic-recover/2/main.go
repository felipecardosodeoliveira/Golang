package main

import (
	"fmt"
	"log"
	"net/http"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entrou no handler")
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	log.Println("Listening on port 8000")
	if err := http.ListenAndServe(":8000", recoverMiddleware(mux)); err != nil {
		log.Fatal("Could not listen : \n", err)
	}
}
