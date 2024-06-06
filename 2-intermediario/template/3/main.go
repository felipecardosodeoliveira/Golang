package main

import (
	"html/template"
	"log"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	c := Cursos{{Nome: "Node JS", CargaHoraria: 44},
		{Nome: "PHP", CargaHoraria: 50},
		{Nome: "Go", CargaHoraria: 40},
	}
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
