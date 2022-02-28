package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "client.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	//tpl = template.Must(template.ParseFiles("templates/client.gohtml"))
	//tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	tpl = template.Must(template.ParseFiles("client.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
