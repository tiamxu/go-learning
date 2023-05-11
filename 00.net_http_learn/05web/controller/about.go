package controller

import (
	"html/template"
	"net/http"
)

func aboutHander(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "about.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func registerAboutRoutes() {
	http.HandleFunc("/about", aboutHander)

}
