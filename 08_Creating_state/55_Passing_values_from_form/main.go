package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text.txt/html; charset=utf-8")
	io.WriteString(w, `
	<!DOCTYPE html>
	<form method="get">
		<input type="text.txt" name="q">
		<input type="submit">
	</form>
	<br>`+v)
}
