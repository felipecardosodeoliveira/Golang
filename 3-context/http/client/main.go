package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8000", nil)
	if err != nil {
		log.Println("Falha na requisição", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Falha na requisição 2", err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	// body, err := io.ReadAll(req.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// print(string(body))
}
