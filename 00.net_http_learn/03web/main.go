package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// t, _ := template.ParseFiles("./index.html")
		// t,_:= template.ParseGlob("*.html")
		t := template.New("index.html")
		t, _ = t.ParseFiles("index.html")
		t.Execute(w, "hello world")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// t, _ := template.ParseFiles("index.html.html")
		// t.Execute(w, "hello world")
		ts, _ := template.ParseFiles("t1.html", "index.html")
		ts.ExecuteTemplate(w, "t1.html", "Hello xxx")
	})

	server := http.Server{Addr: ":9091", Handler: nil}
	server.ListenAndServe()
}
