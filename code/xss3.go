package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("param1")
	tmp1 := template.New("hello")
	tmp1, _ = tmp1.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmp1.ExecuteTemplate(w, "T", param1)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
