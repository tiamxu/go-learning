package controller

import (
	"html/template"
	"net/http"
)

func homeHander(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "home.html")
	t.ExecuteTemplate(w, "layout", "Hello world")
}

func registerHomeRoutes() {
	http.HandleFunc("/home", homeHander)

}
