package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Clayton", "clayton.santos@gmail.com",}		

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escuntando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}