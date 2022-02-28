package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  56,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "client.gohtml", p1)
	if err != nil {
		log.Fatal(err)
	}
}
