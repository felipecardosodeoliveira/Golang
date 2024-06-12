package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser
	jsonData := `{"foo":"bar", "num":123, "bool":true, "arr":[1,2]}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("num=%d\n", v.GetInt("num"))
	fmt.Printf("bool=%t\n", v.GetBool("bool"))
	fmt.Printf("arr=%v\n", v.GetArray("arr"))

	fmt.Println("*********************************************************")

	jsonUser := `{"user": {"name": "Felipe", "age": 30}}`

	v2, err := p.Parse(jsonUser)
	if err != nil {
		panic(err)
	}
	user := v2.GetObject("user")
	fmt.Printf("User name %s \n", user.Get("name"))
}
