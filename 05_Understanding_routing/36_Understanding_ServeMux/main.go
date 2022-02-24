package main

import (
	"io"
	"log"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy doggy")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty kitty")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
