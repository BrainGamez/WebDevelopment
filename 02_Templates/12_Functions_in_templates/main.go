package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s

}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
