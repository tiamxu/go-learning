package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("templates")
		// t, err := t.ParseGlob("./template/*.html")
		// template.Must(t, err)
		template.Must(t.ParseGlob("./template/*.html"))
		// w.Write([]byte("hello..."))
		// log.Println(r.URL)
		// fmt.Fprintln(w, r.URL.Path)
		fileName := r.URL.Path[1:]
		t = t.Lookup(fileName)
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/css/", http.FileServer(http.Dir("wwwroot")))

	http.ListenAndServe(":8090", nil)
}
