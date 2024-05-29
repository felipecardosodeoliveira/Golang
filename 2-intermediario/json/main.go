package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Name string `json:"n"`
	Age  int    `json:"a"`
}

func main() {
	p := Pessoa{Name: "Felipe", Age: 32}
	res, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
	err = json.NewEncoder(os.Stdout).Encode(p)
	if err != nil {
		panic(err)
	}

	jp := []byte(`{"n": "Felipe", "a": 32}`)
	var pX Pessoa
	err = json.Unmarshal(jp, &pX)
	if err != nil {
		panic(err)
	}

	fmt.Println(pX.Age)
}
