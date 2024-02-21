package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type users struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	user := users{
		"Eduardo",
		"Eduardo@gmail.com",
	}
	templates.ExecuteTemplate(w, "home.html", user)
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Printf("Ouindo na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
