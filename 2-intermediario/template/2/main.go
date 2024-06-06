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

func main() {
	c := Curso{Nome: "Node JS", CargaHoraria: 44}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} carga horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
