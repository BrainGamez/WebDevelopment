package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Nobody knows")
	if err != nil {
		log.Fatal(err)
	}
}
