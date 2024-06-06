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
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso {{.Nome}} carga horária {{.CargaHoraria}}")
	err := tmp.Execute(os.Stdout, c)
	if err != nil {
		log.Fatal(err)
	}
}
