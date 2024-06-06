package main

import (
	"html/template"
	"log"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := Cursos{{Nome: "Node JS", CargaHoraria: 44},
			{Nome: "PHP", CargaHoraria: 50},
			{Nome: "Go", CargaHoraria: 40},
		}
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, c)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8000", nil)
}
