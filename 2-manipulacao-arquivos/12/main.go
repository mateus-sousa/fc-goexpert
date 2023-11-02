package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Course struct {
	Name     string
	WorkLoad int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(w, Courses{
			{Name: "Go", WorkLoad: 40},
			{Name: "Docker", WorkLoad: 30},
			{Name: "Kubernetes", WorkLoad: 50},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", mux)
}
