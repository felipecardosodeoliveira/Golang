package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	jsonVar := bytes.NewBuffer([]byte(`{Name: Felipe, Idade: 30}`))
	resp, err := c.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
