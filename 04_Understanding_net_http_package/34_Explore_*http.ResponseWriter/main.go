package main

import (
	"fmt"
	"log"
	"net/http"
)

type attachedType int

func (m attachedType) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Bryan-Key", "This is from Bryan")
	w.Header().Set("Content-Type", "text.txt/html; charset=utf-8")
	_, err := fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var d attachedType
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln(err)
	}
}
