package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

/*
*
identifica alterações com o sistema em execução
*
*/

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()
	MarshalConfig("config.json")

	done := make(chan bool) // esse canal está segurando o programa e aguardando alteraçõs em configs
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("Event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
					fmt.Println("Modified file: ", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error: ", err)
			}
		}
	}()
	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
