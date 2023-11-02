package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloWorld)
	mux.Handle("/blog", blog{})
	http.ListenAndServe(":8080", mux)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD"))
}

type blog struct {
	typeBlog string
}

// O nome do metodo TEM QUE SER ServeHTTP
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get blog"))
}
