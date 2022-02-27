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

type sage struct {
	Name  string
	Motto string
}

func main() {
	//sages := []string{"Gandhi", "MLK", "Buddha", "Jesus"}

	Buddha := sage{
		Name:  "Buddha",
		Motto: "Belief of no beliefs",
	}

	Gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	MLK := sage{
		Name:  "Martin Luther King",
		Motto: "I have a dream",
	}

	sages := []sage{Buddha, Gandhi, MLK}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
