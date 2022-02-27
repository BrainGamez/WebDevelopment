package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := struct {
		Words []string
		Lname string
	}{
		xs,
		"Francisca",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
