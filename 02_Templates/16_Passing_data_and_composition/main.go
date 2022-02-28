package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"CSCI-40", "Introduction to programming in Go", "4"},
				course{"CSCI-130", "Introduction  programming with Go", "4"},
				course{"CSCI-140", "Mobile Apps using Go", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"CSCI-50", "Advanced programming in Go", "8"},
				course{"CSCI-190", "Advanced programming with Go", "8"},
				course{"CSCI-191", "Advanced Apps using Go", "8"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "client.gohtml", y)
	if err != nil {
		log.Fatal(err)
	}
}
