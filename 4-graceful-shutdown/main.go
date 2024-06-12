package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World \n"))
	})

	go func() {
		fmt.Println("server is running")
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not liste on %s %v", server.Addr, err)
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fmt.Println("Shutdown Server")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Server stopped")
}
