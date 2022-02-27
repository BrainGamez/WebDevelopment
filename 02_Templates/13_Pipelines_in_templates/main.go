package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fDayMontYear": dayMonthYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
