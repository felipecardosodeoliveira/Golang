package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Multiply(n int) int {
	return n * 2
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := Cursos{{Nome: "Node JS", CargaHoraria: 4},
			{Nome: "PHP", CargaHoraria: 5},
			{Nome: "Go", CargaHoraria: 4},
			{Nome: "Java", CargaHoraria: 10},
		}

		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper, "Multiply": Multiply})
		t = template.Must(t.ParseFiles(templates...))

		// t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, c)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.ListenAndServe(":8000", nil)
}
