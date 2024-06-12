package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com suceso \n"))
		return
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
		return
	}
}
